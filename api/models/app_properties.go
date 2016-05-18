package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*AppProperties app properties

swagger:model AppProperties
*/
type AppProperties struct {

	/* id
	 */
	ID int32 `json:"id,omitempty"`

	/* name
	 */
	Name string `json:"name,omitempty"`

	/* replicas
	 */
	Replicas *int32 `json:"replicas,omitempty"`
}

// Validate validates this app properties
func (m *AppProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
