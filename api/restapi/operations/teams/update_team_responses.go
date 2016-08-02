package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/teresa/api/models"
)

/*UpdateTeamOK Team details

swagger:response updateTeamOK
*/
type UpdateTeamOK struct {

	// In: body
	Payload *models.Team `json:"body,omitempty"`
}

// NewUpdateTeamOK creates UpdateTeamOK with default headers values
func NewUpdateTeamOK() *UpdateTeamOK {
	return &UpdateTeamOK{}
}

// WithPayload adds the payload to the update team o k response
func (o *UpdateTeamOK) WithPayload(payload *models.Team) *UpdateTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update team o k response
func (o *UpdateTeamOK) SetPayload(payload *models.Team) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateTeamDefault Error

swagger:response updateTeamDefault
*/
type UpdateTeamDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateTeamDefault creates UpdateTeamDefault with default headers values
func NewUpdateTeamDefault(code int) *UpdateTeamDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateTeamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update team default response
func (o *UpdateTeamDefault) WithStatusCode(code int) *UpdateTeamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update team default response
func (o *UpdateTeamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update team default response
func (o *UpdateTeamDefault) WithPayload(payload *models.Error) *UpdateTeamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update team default response
func (o *UpdateTeamDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTeamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
