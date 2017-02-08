package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/teresa-api/models"
)

// NewPartialUpdateAppParams creates a new PartialUpdateAppParams object
// with the default values initialized.
func NewPartialUpdateAppParams() *PartialUpdateAppParams {
	var ()
	return &PartialUpdateAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPartialUpdateAppParamsWithTimeout creates a new PartialUpdateAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPartialUpdateAppParamsWithTimeout(timeout time.Duration) *PartialUpdateAppParams {
	var ()
	return &PartialUpdateAppParams{

		timeout: timeout,
	}
}

/*PartialUpdateAppParams contains all the parameters to send to the API endpoint
for the partial update app operation typically these are written to a http.Request
*/
type PartialUpdateAppParams struct {

	/*AppName
	  App name

	*/
	AppName string
	/*Body*/
	Body []*models.PatchAppRequest

	timeout time.Duration
}

// WithAppName adds the appName to the partial update app params
func (o *PartialUpdateAppParams) WithAppName(appName string) *PartialUpdateAppParams {
	o.AppName = appName
	return o
}

// WithBody adds the body to the partial update app params
func (o *PartialUpdateAppParams) WithBody(body []*models.PatchAppRequest) *PartialUpdateAppParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PartialUpdateAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
