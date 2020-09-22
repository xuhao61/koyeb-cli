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

	strfmt "github.com/go-openapi/strfmt"
)

// NewCatalogListCatalogStoresParams creates a new CatalogListCatalogStoresParams object
// with the default values initialized.
func NewCatalogListCatalogStoresParams() *CatalogListCatalogStoresParams {
	var ()
	return &CatalogListCatalogStoresParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogListCatalogStoresParamsWithTimeout creates a new CatalogListCatalogStoresParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogListCatalogStoresParamsWithTimeout(timeout time.Duration) *CatalogListCatalogStoresParams {
	var ()
	return &CatalogListCatalogStoresParams{

		timeout: timeout,
	}
}

// NewCatalogListCatalogStoresParamsWithContext creates a new CatalogListCatalogStoresParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogListCatalogStoresParamsWithContext(ctx context.Context) *CatalogListCatalogStoresParams {
	var ()
	return &CatalogListCatalogStoresParams{

		Context: ctx,
	}
}

// NewCatalogListCatalogStoresParamsWithHTTPClient creates a new CatalogListCatalogStoresParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogListCatalogStoresParamsWithHTTPClient(client *http.Client) *CatalogListCatalogStoresParams {
	var ()
	return &CatalogListCatalogStoresParams{
		HTTPClient: client,
	}
}

/*CatalogListCatalogStoresParams contains all the parameters to send to the API endpoint
for the catalog list catalog stores operation typically these are written to a http.Request
*/
type CatalogListCatalogStoresParams struct {

	/*Limit*/
	Limit *string
	/*Name*/
	Name *string
	/*Offset*/
	Offset *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) WithTimeout(timeout time.Duration) *CatalogListCatalogStoresParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) WithContext(ctx context.Context) *CatalogListCatalogStoresParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) WithHTTPClient(client *http.Client) *CatalogListCatalogStoresParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) WithLimit(limit *string) *CatalogListCatalogStoresParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithName adds the name to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) WithName(name *string) *CatalogListCatalogStoresParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) WithOffset(offset *string) *CatalogListCatalogStoresParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the catalog list catalog stores params
func (o *CatalogListCatalogStoresParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogListCatalogStoresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}