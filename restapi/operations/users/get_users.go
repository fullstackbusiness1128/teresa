package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"github.com/luizalabs/teresa-api/models"
)

// GetUsersHandlerFunc turns a function with the right signature into a get users handler
type GetUsersHandlerFunc func(GetUsersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersHandlerFunc) Handle(params GetUsersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetUsersHandler interface for that can handle valid get users params
type GetUsersHandler interface {
	Handle(GetUsersParams, interface{}) middleware.Responder
}

// NewGetUsers creates a new http.Handler for the get users operation
func NewGetUsers(ctx *middleware.Context, handler GetUsersHandler) *GetUsers {
	return &GetUsers{Context: ctx, Handler: handler}
}

/*GetUsers swagger:route GET /users users getUsers

Get users

Find and filter users

*/
type GetUsers struct {
	Context *middleware.Context
	Handler GetUsersHandler
}

func (o *GetUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetUsersParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

/*GetUsersOKBodyBody get users o k body body

swagger:model GetUsersOKBodyBody
*/
type GetUsersOKBodyBody struct {
	models.Pagination

	/* items

	Required: true
	*/
	Items []*models.User `json:"items"`
}

// Validate validates this get users o k body body
func (o *GetUsersOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.Pagination.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersOKBodyBody) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("getUsersOK"+"."+"items", "body", o.Items); err != nil {
		return err
	}

	for i := 0; i < len(o.Items); i++ {

		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {

			if err := o.Items[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
