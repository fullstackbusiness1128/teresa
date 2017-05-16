package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new teams API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for teams API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddUserToTeam adds user to team
*/
func (a *Client) AddUserToTeam(params *AddUserToTeamParams, authInfo runtime.ClientAuthInfoWriter) (*AddUserToTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUserToTeamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addUserToTeam",
		Method:             "POST",
		PathPattern:        "/teams/{team_name}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddUserToTeamReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddUserToTeamOK), nil
}

/*
DeleteTeam deletes team

Delete team
*/
func (a *Client) DeleteTeam(params *DeleteTeamParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTeamNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTeamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTeam",
		Method:             "DELETE",
		PathPattern:        "/teams/{team_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTeamReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteTeamNoContent), nil
}

/*
GetTeamDetail gets team details

Get the details of a specific team
*/
func (a *Client) GetTeamDetail(params *GetTeamDetailParams, authInfo runtime.ClientAuthInfoWriter) (*GetTeamDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTeamDetailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTeamDetail",
		Method:             "GET",
		PathPattern:        "/teams/{team_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTeamDetailReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTeamDetailOK), nil
}

/*
GetTeams gets teams

Get a list of teams
*/
func (a *Client) GetTeams(params *GetTeamsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTeamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTeamsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTeams",
		Method:             "GET",
		PathPattern:        "/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTeamsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTeamsOK), nil
}

/*
UpdateTeam updates team

Update team details
*/
func (a *Client) UpdateTeam(params *UpdateTeamParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTeamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateTeam",
		Method:             "PUT",
		PathPattern:        "/teams/{team_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTeamReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateTeamOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
