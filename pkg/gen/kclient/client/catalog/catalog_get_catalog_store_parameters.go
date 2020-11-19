// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewCatalogGetCatalogStoreParams creates a new CatalogGetCatalogStoreParams object
// with the default values initialized.
func NewCatalogGetCatalogStoreParams() *CatalogGetCatalogStoreParams {
	var ()
	return &CatalogGetCatalogStoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogGetCatalogStoreParamsWithTimeout creates a new CatalogGetCatalogStoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogGetCatalogStoreParamsWithTimeout(timeout time.Duration) *CatalogGetCatalogStoreParams {
	var ()
	return &CatalogGetCatalogStoreParams{

		timeout: timeout,
	}
}

// NewCatalogGetCatalogStoreParamsWithContext creates a new CatalogGetCatalogStoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogGetCatalogStoreParamsWithContext(ctx context.Context) *CatalogGetCatalogStoreParams {
	var ()
	return &CatalogGetCatalogStoreParams{

		Context: ctx,
	}
}

// NewCatalogGetCatalogStoreParamsWithHTTPClient creates a new CatalogGetCatalogStoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogGetCatalogStoreParamsWithHTTPClient(client *http.Client) *CatalogGetCatalogStoreParams {
	var ()
	return &CatalogGetCatalogStoreParams{
		HTTPClient: client,
	}
}

/*CatalogGetCatalogStoreParams contains all the parameters to send to the API endpoint
for the catalog get catalog store operation typically these are written to a http.Request
*/
type CatalogGetCatalogStoreParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog get catalog store params
func (o *CatalogGetCatalogStoreParams) WithTimeout(timeout time.Duration) *CatalogGetCatalogStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog get catalog store params
func (o *CatalogGetCatalogStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog get catalog store params
func (o *CatalogGetCatalogStoreParams) WithContext(ctx context.Context) *CatalogGetCatalogStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog get catalog store params
func (o *CatalogGetCatalogStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog get catalog store params
func (o *CatalogGetCatalogStoreParams) WithHTTPClient(client *http.Client) *CatalogGetCatalogStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog get catalog store params
func (o *CatalogGetCatalogStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the catalog get catalog store params
func (o *CatalogGetCatalogStoreParams) WithName(name string) *CatalogGetCatalogStoreParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the catalog get catalog store params
func (o *CatalogGetCatalogStoreParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogGetCatalogStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}