package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*App app

swagger:model App
*/
type App struct {

	/* list of the external addresses of the app
	 */
	AddressList []string `json:"addressList,omitempty"`

	/* creator
	 */
	Creator *User `json:"creator,omitempty"`

	/* deployment list
	 */
	DeploymentList []*Deployment `json:"deploymentList,omitempty"`

	/* env vars
	 */
	EnvVars []*EnvVar `json:"envVars,omitempty"`

	/* id
	 */
	ID int64 `json:"id,omitempty"`

	/* name

	Required: true
	*/
	Name *string `json:"name"`

	/* number of PODs running the app
	 */
	Scale *int64 `json:"scale,omitempty"`
}

// Validate validates this app
func (m *App) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreator(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDeploymentList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnvVars(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *App) validateAddressList(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressList) { // not required
		return nil
	}

	return nil
}

func (m *App) validateCreator(formats strfmt.Registry) error {

	if swag.IsZero(m.Creator) { // not required
		return nil
	}

	if m.Creator != nil {

		if err := m.Creator.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *App) validateDeploymentList(formats strfmt.Registry) error {

	if swag.IsZero(m.DeploymentList) { // not required
		return nil
	}

	for i := 0; i < len(m.DeploymentList); i++ {

		if swag.IsZero(m.DeploymentList[i]) { // not required
			continue
		}

		if m.DeploymentList[i] != nil {

			if err := m.DeploymentList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *App) validateEnvVars(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvVars) { // not required
		return nil
	}

	for i := 0; i < len(m.EnvVars); i++ {

		if swag.IsZero(m.EnvVars[i]) { // not required
			continue
		}

		if m.EnvVars[i] != nil {

			if err := m.EnvVars[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *App) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
