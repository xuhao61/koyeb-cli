// Code generated by go-swagger; DO NOT EDIT.

package stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new stack API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stack API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	StackDeleteStack(params *StackDeleteStackParams, authInfo runtime.ClientAuthInfoWriter) (*StackDeleteStackOK, error)

	StackGetStack(params *StackGetStackParams, authInfo runtime.ClientAuthInfoWriter) (*StackGetStackOK, error)

	StackGetStackRevision(params *StackGetStackRevisionParams, authInfo runtime.ClientAuthInfoWriter) (*StackGetStackRevisionOK, error)

	StackListStackRevisions(params *StackListStackRevisionsParams, authInfo runtime.ClientAuthInfoWriter) (*StackListStackRevisionsOK, error)

	StackListStacks(params *StackListStacksParams, authInfo runtime.ClientAuthInfoWriter) (*StackListStacksOK, error)

	StackNewStack(params *StackNewStackParams, authInfo runtime.ClientAuthInfoWriter) (*StackNewStackOK, error)

	StackUpdateStack(params *StackUpdateStackParams, authInfo runtime.ClientAuthInfoWriter) (*StackUpdateStackOK, error)

	StackUpdateStack2(params *StackUpdateStack2Params, authInfo runtime.ClientAuthInfoWriter) (*StackUpdateStack2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  StackDeleteStack deletes an existing stack
*/
func (a *Client) StackDeleteStack(params *StackDeleteStackParams, authInfo runtime.ClientAuthInfoWriter) (*StackDeleteStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStackDeleteStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stack_DeleteStack",
		Method:             "DELETE",
		PathPattern:        "/v1/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StackDeleteStackReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StackDeleteStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StackDeleteStackDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StackGetStack fetches a stack
*/
func (a *Client) StackGetStack(params *StackGetStackParams, authInfo runtime.ClientAuthInfoWriter) (*StackGetStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStackGetStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stack_GetStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StackGetStackReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StackGetStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StackGetStackDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StackGetStackRevision fetches a stack revision
*/
func (a *Client) StackGetStackRevision(params *StackGetStackRevisionParams, authInfo runtime.ClientAuthInfoWriter) (*StackGetStackRevisionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStackGetStackRevisionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stack_GetStackRevision",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{stack_id}/revisions/{sha}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StackGetStackRevisionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StackGetStackRevisionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StackGetStackRevisionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StackListStackRevisions lists stack revisions
*/
func (a *Client) StackListStackRevisions(params *StackListStackRevisionsParams, authInfo runtime.ClientAuthInfoWriter) (*StackListStackRevisionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStackListStackRevisionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stack_ListStackRevisions",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{stack_id}/revisions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StackListStackRevisionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StackListStackRevisionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StackListStackRevisionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StackListStacks lists stacks
*/
func (a *Client) StackListStacks(params *StackListStacksParams, authInfo runtime.ClientAuthInfoWriter) (*StackListStacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStackListStacksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stack_ListStacks",
		Method:             "GET",
		PathPattern:        "/v1/stacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StackListStacksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StackListStacksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StackListStacksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StackNewStack creates a stack
*/
func (a *Client) StackNewStack(params *StackNewStackParams, authInfo runtime.ClientAuthInfoWriter) (*StackNewStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStackNewStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stack_NewStack",
		Method:             "POST",
		PathPattern:        "/v1/stacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StackNewStackReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StackNewStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StackNewStackDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StackUpdateStack updates an existing stack
*/
func (a *Client) StackUpdateStack(params *StackUpdateStackParams, authInfo runtime.ClientAuthInfoWriter) (*StackUpdateStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStackUpdateStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stack_UpdateStack",
		Method:             "PUT",
		PathPattern:        "/v1/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StackUpdateStackReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StackUpdateStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StackUpdateStackDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StackUpdateStack2 updates an existing stack
*/
func (a *Client) StackUpdateStack2(params *StackUpdateStack2Params, authInfo runtime.ClientAuthInfoWriter) (*StackUpdateStack2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStackUpdateStack2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stack_UpdateStack2",
		Method:             "PATCH",
		PathPattern:        "/v1/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StackUpdateStack2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StackUpdateStack2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StackUpdateStack2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
