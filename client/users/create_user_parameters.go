package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/tapi/models"
)

// NewCreateUserParams creates a new CreateUserParams object
// with the default values initialized.
func NewCreateUserParams() *CreateUserParams {
	var ()
	return &CreateUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserParamsWithTimeout creates a new CreateUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUserParamsWithTimeout(timeout time.Duration) *CreateUserParams {
	var ()
	return &CreateUserParams{

		timeout: timeout,
	}
}

/*CreateUserParams contains all the parameters to send to the API endpoint
for the create user operation typically these are written to a http.Request
*/
type CreateUserParams struct {

	/*Body*/
	Body *models.User

	timeout time.Duration
}

// WithBody adds the body to the create user params
func (o *CreateUserParams) WithBody(body *models.User) *CreateUserParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.User)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
