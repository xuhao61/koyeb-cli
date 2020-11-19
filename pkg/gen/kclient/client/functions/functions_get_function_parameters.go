// Code generated by go-swagger; DO NOT EDIT.

package functions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewFunctionsGetFunctionParams creates a new FunctionsGetFunctionParams object
// with the default values initialized.
func NewFunctionsGetFunctionParams() *FunctionsGetFunctionParams {
	var ()
	return &FunctionsGetFunctionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFunctionsGetFunctionParamsWithTimeout creates a new FunctionsGetFunctionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFunctionsGetFunctionParamsWithTimeout(timeout time.Duration) *FunctionsGetFunctionParams {
	var ()
	return &FunctionsGetFunctionParams{

		timeout: timeout,
	}
}

// NewFunctionsGetFunctionParamsWithContext creates a new FunctionsGetFunctionParams object
// with the default values initialized, and the ability to set a context for a request
func NewFunctionsGetFunctionParamsWithContext(ctx context.Context) *FunctionsGetFunctionParams {
	var ()
	return &FunctionsGetFunctionParams{

		Context: ctx,
	}
}

// NewFunctionsGetFunctionParamsWithHTTPClient creates a new FunctionsGetFunctionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFunctionsGetFunctionParamsWithHTTPClient(client *http.Client) *FunctionsGetFunctionParams {
	var ()
	return &FunctionsGetFunctionParams{
		HTTPClient: client,
	}
}

/*FunctionsGetFunctionParams contains all the parameters to send to the API endpoint
for the functions get function operation typically these are written to a http.Request
*/
type FunctionsGetFunctionParams struct {

	/*Function*/
	Function string
	/*Sha
	  The sha of the revision (either short of long form, `:latest` returns the latest deployed revision)

	*/
	Sha string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the functions get function params
func (o *FunctionsGetFunctionParams) WithTimeout(timeout time.Duration) *FunctionsGetFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the functions get function params
func (o *FunctionsGetFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the functions get function params
func (o *FunctionsGetFunctionParams) WithContext(ctx context.Context) *FunctionsGetFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the functions get function params
func (o *FunctionsGetFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the functions get function params
func (o *FunctionsGetFunctionParams) WithHTTPClient(client *http.Client) *FunctionsGetFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the functions get function params
func (o *FunctionsGetFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFunction adds the function to the functions get function params
func (o *FunctionsGetFunctionParams) WithFunction(function string) *FunctionsGetFunctionParams {
	o.SetFunction(function)
	return o
}

// SetFunction adds the function to the functions get function params
func (o *FunctionsGetFunctionParams) SetFunction(function string) {
	o.Function = function
}

// WithSha adds the sha to the functions get function params
func (o *FunctionsGetFunctionParams) WithSha(sha string) *FunctionsGetFunctionParams {
	o.SetSha(sha)
	return o
}

// SetSha adds the sha to the functions get function params
func (o *FunctionsGetFunctionParams) SetSha(sha string) {
	o.Sha = sha
}

// WithStackID adds the stackID to the functions get function params
func (o *FunctionsGetFunctionParams) WithStackID(stackID string) *FunctionsGetFunctionParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the functions get function params
func (o *FunctionsGetFunctionParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *FunctionsGetFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param function
	if err := r.SetPathParam("function", o.Function); err != nil {
		return err
	}

	// path param sha
	if err := r.SetPathParam("sha", o.Sha); err != nil {
		return err
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}