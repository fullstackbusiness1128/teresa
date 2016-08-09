package restapi

import (
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	flags "github.com/jessevdk/go-flags"
	graceful "github.com/tylerb/graceful"

	"github.com/luizalabs/tapi/restapi/operations"
)

//go:generate swagger generate server -t ../.. -A Teresa -f ./swagger.yml

// NewServer creates a new api teresa server but does not configure it
func NewServer(api *operations.TeresaAPI) *Server {
	s := new(Server)
	s.api = api
	return s
}

// ConfigureAPI configures the API and handlers. Needs to be called before Serve
func (s *Server) ConfigureAPI() {
	if s.api != nil {
		s.handler = configureAPI(s.api)
	}
}

// ConfigureFlags configures the additional flags defined by the handlers. Needs to be called before the parser.Parse
func (s *Server) ConfigureFlags() {
	if s.api != nil {
		configureFlags(s.api)
	}
}

// Server for the teresa API
type Server struct {
	ForcedSchemes []string `long:"override-scheme" choice:"http" choice:"https" choice:"wss" description:"Override schemes defined in the swagger spec."`

	SocketPath flags.Filename `long:"unix-socket" description:"the unix socket to listen on"`
	HTTPServer string         `long:"http-server" description:"Host:Port for HTTP Server"`

	HTTPSServer string         `long:"https-server" description:"Host:Port for HTTPS Server"`
	HTTPSCert   flags.Filename `long:"https-tls-cert" description:"the certificate to use for secure connections"`
	HTTPSKey    flags.Filename `long:"https-tls-key" description:"the private key to use for secure connections"`

	domainSocketL net.Listener
	httpsServerL  net.Listener
	httpServerL   net.Listener

	api          *operations.TeresaAPI
	handler      http.Handler
	hasListeners bool
}

// Logf logs message either via defined user logger or via system one if no user logger is defined.
func (s *Server) Logf(f string, args ...interface{}) {
	if s.api != nil && s.api.Logger != nil {
		s.api.Logger(f, args...)
	} else {
		log.Printf(f, args...)
	}
}

// SetAPI configures the server with the specified API. Needs to be called before Serve
func (s *Server) SetAPI(api *operations.TeresaAPI) {
	if api == nil {
		s.api = nil
		s.handler = nil
		return
	}

	s.api = api
	s.api.Logger = log.Printf
	s.handler = configureAPI(api)
}

// Serve the api
func (s *Server) Serve() error {
	var wg sync.WaitGroup

	if s.HTTPServer == "" && s.HTTPSServer == "" && s.SocketPath == "" {
		return errors.New("At least one listening server have to be defined")
	}

	if s.HTTPServer != "" {
		listener, err := net.Listen("tcp", s.HTTPServer)
		if err != nil {
			return err
		}

		s.httpServerL = listener

		httpServer := &graceful.Server{Server: new(http.Server)}
		httpServer.Handler = s.handler
		wg.Add(1)
		go func(l net.Listener) {
			defer wg.Done()
			s.Logf("Serving rdb at http://%s", s.httpServerL.Addr())
			if err := httpServer.Serve(tcpKeepAliveListener{l.(*net.TCPListener)}); err != nil {
				log.Fatalln(err)
			}
		}(s.httpServerL)
	}

	if s.HTTPSServer != "" {
		tlsListener, err := net.Listen("tcp", s.HTTPSServer)
		if err != nil {
			return err
		}
		s.httpsServerL = tlsListener

		if s.HTTPSCert == "" {
			return errors.New("TLS Certificate is not provided for HTTPS")
		}
		if s.HTTPSKey == "" {
			return errors.New("TLS Key is not provided for HTTPS")
		}
		httpsServer := &graceful.Server{Server: new(http.Server)}
		httpsServer.Handler = s.handler
		httpsServer.TLSConfig = new(tls.Config)
		httpsServer.TLSConfig.NextProtos = []string{"http/1.1"}

		// https://www.owasp.org/index.php/Transport_Layer_Protection_Cheat_Sheet#Rule_-_Only_Support_Strong_Protocols
		httpsServer.TLSConfig.MinVersion = tls.VersionTLS12
		httpsServer.TLSConfig.Certificates = make([]tls.Certificate, 1)
		httpsServer.TLSConfig.Certificates[0], err = tls.LoadX509KeyPair(string(s.HTTPSCert), string(s.HTTPSKey))

		configureTLS(httpsServer.TLSConfig)
		wg.Add(1)
		go func(l net.Listener) {
			defer wg.Done()
			s.Logf("Serving rdb at http://%s", s.httpsServerL.Addr())
			if err := httpsServer.Serve(tcpKeepAliveListener{l.(*net.TCPListener)}); err != nil {
				log.Fatalln(err)
			}
		}(s.httpsServerL)
	}

	if s.SocketPath != "" {
		domSockListener, err := net.Listen("unix", string(s.SocketPath))
		if err != nil {
			return err
		}
		s.domainSocketL = domSockListener

		domainSocket := &graceful.Server{Server: new(http.Server)}
		domainSocket.Handler = s.handler
		wg.Add(1)
		go func(l net.Listener) {
			defer wg.Done()
			s.Logf("Serving rdb at unix://%s", s.SocketPath)
			if err := domainSocket.Serve(l); err != nil {
				log.Fatalln(err)
			}
		}(s.domainSocketL)
	}

	wg.Wait()
	return nil
}

// Shutdown server and clean up resources
func (s *Server) Shutdown() error {
	s.api.ServerShutdown()
	return nil
}

// GetHandler returns a handler useful for testing
func (s *Server) GetHandler() http.Handler {
	return s.handler
}

// tcpKeepAliveListener is copied from the stdlib net/http package

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}
