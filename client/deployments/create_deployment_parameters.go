package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"os"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateDeploymentParams creates a new CreateDeploymentParams object
// with the default values initialized.
func NewCreateDeploymentParams() *CreateDeploymentParams {
	var ()
	return &CreateDeploymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeploymentParamsWithTimeout creates a new CreateDeploymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDeploymentParamsWithTimeout(timeout time.Duration) *CreateDeploymentParams {
	var ()
	return &CreateDeploymentParams{

		timeout: timeout,
	}
}

/*CreateDeploymentParams contains all the parameters to send to the API endpoint
for the create deployment operation typically these are written to a http.Request
*/
type CreateDeploymentParams struct {

	/*AppTarball*/
	AppTarball os.File
	/*AppName
	  App name

	*/
	AppName string
	/*Description*/
	Description *string

	timeout time.Duration
}

// WithAppTarball adds the appTarball to the create deployment params
func (o *CreateDeploymentParams) WithAppTarball(appTarball os.File) *CreateDeploymentParams {
	o.AppTarball = appTarball
	return o
}

// WithAppName adds the appName to the create deployment params
func (o *CreateDeploymentParams) WithAppName(appName string) *CreateDeploymentParams {
	o.AppName = appName
	return o
}

// WithDescription adds the description to the create deployment params
func (o *CreateDeploymentParams) WithDescription(description *string) *CreateDeploymentParams {
	o.Description = description
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// form file param appTarball
	if err := r.SetFileParam("appTarball", &o.AppTarball); err != nil {
		return err
	}

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
