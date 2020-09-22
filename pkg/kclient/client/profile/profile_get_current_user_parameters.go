// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewProfileGetCurrentUserParams creates a new ProfileGetCurrentUserParams object
// with the default values initialized.
func NewProfileGetCurrentUserParams() *ProfileGetCurrentUserParams {

	return &ProfileGetCurrentUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProfileGetCurrentUserParamsWithTimeout creates a new ProfileGetCurrentUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProfileGetCurrentUserParamsWithTimeout(timeout time.Duration) *ProfileGetCurrentUserParams {

	return &ProfileGetCurrentUserParams{

		timeout: timeout,
	}
}

// NewProfileGetCurrentUserParamsWithContext creates a new ProfileGetCurrentUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewProfileGetCurrentUserParamsWithContext(ctx context.Context) *ProfileGetCurrentUserParams {

	return &ProfileGetCurrentUserParams{

		Context: ctx,
	}
}

// NewProfileGetCurrentUserParamsWithHTTPClient creates a new ProfileGetCurrentUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProfileGetCurrentUserParamsWithHTTPClient(client *http.Client) *ProfileGetCurrentUserParams {

	return &ProfileGetCurrentUserParams{
		HTTPClient: client,
	}
}

/*ProfileGetCurrentUserParams contains all the parameters to send to the API endpoint
for the profile get current user operation typically these are written to a http.Request
*/
type ProfileGetCurrentUserParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the profile get current user params
func (o *ProfileGetCurrentUserParams) WithTimeout(timeout time.Duration) *ProfileGetCurrentUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the profile get current user params
func (o *ProfileGetCurrentUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the profile get current user params
func (o *ProfileGetCurrentUserParams) WithContext(ctx context.Context) *ProfileGetCurrentUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the profile get current user params
func (o *ProfileGetCurrentUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the profile get current user params
func (o *ProfileGetCurrentUserParams) WithHTTPClient(client *http.Client) *ProfileGetCurrentUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the profile get current user params
func (o *ProfileGetCurrentUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ProfileGetCurrentUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}