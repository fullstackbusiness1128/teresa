package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	loads "github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/luizalabs/tapi/restapi"
	"github.com/luizalabs/tapi/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTeresaAPI()
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = `Teresa API`
	parser.LongDescription = `The Teresa PaaS API`

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}

	if len(server.ForcedSchemes) > 0 {
		d := json.NewDecoder(bytes.NewReader(swaggerSpec.Raw()))
		d.UseNumber()

		var data interface{}
		if err := d.Decode(&data); err != nil {
			panic("Broken schema!")
		}
		mdata, _ := data.(map[string]interface{})
		mdata["schemes"] = server.ForcedSchemes

		rawSpec, _ := json.MarshalIndent(data, " ", "  ")
		swaggerSpec, err = loads.Analyzed(rawSpec, "")
		if err != nil {
			log.Fatalln(err)
		}
	}

	api.SetSpec(swaggerSpec)
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		server.Shutdown()
		log.Fatalln(err)
	}
}
