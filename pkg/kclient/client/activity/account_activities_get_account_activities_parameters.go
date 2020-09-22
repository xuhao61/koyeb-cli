// Code generated by go-swagger; DO NOT EDIT.

package activity

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

// NewAccountActivitiesGetAccountActivitiesParams creates a new AccountActivitiesGetAccountActivitiesParams object
// with the default values initialized.
func NewAccountActivitiesGetAccountActivitiesParams() *AccountActivitiesGetAccountActivitiesParams {
	var ()
	return &AccountActivitiesGetAccountActivitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAccountActivitiesGetAccountActivitiesParamsWithTimeout creates a new AccountActivitiesGetAccountActivitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountActivitiesGetAccountActivitiesParamsWithTimeout(timeout time.Duration) *AccountActivitiesGetAccountActivitiesParams {
	var ()
	return &AccountActivitiesGetAccountActivitiesParams{

		timeout: timeout,
	}
}

// NewAccountActivitiesGetAccountActivitiesParamsWithContext creates a new AccountActivitiesGetAccountActivitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountActivitiesGetAccountActivitiesParamsWithContext(ctx context.Context) *AccountActivitiesGetAccountActivitiesParams {
	var ()
	return &AccountActivitiesGetAccountActivitiesParams{

		Context: ctx,
	}
}

// NewAccountActivitiesGetAccountActivitiesParamsWithHTTPClient creates a new AccountActivitiesGetAccountActivitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAccountActivitiesGetAccountActivitiesParamsWithHTTPClient(client *http.Client) *AccountActivitiesGetAccountActivitiesParams {
	var ()
	return &AccountActivitiesGetAccountActivitiesParams{
		HTTPClient: client,
	}
}

/*AccountActivitiesGetAccountActivitiesParams contains all the parameters to send to the API endpoint
for the account activities get account activities operation typically these are written to a http.Request
*/
type AccountActivitiesGetAccountActivitiesParams struct {

	/*Limit*/
	Limit *string
	/*Offset*/
	Offset *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the account activities get account activities params
func (o *AccountActivitiesGetAccountActivitiesParams) WithTimeout(timeout time.Duration) *AccountActivitiesGetAccountActivitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account activities get account activities params
func (o *AccountActivitiesGetAccountActivitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account activities get account activities params
func (o *AccountActivitiesGetAccountActivitiesParams) WithContext(ctx context.Context) *AccountActivitiesGetAccountActivitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account activities get account activities params
func (o *AccountActivitiesGetAccountActivitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account activities get account activities params
func (o *AccountActivitiesGetAccountActivitiesParams) WithHTTPClient(client *http.Client) *AccountActivitiesGetAccountActivitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account activities get account activities params
func (o *AccountActivitiesGetAccountActivitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the account activities get account activities params
func (o *AccountActivitiesGetAccountActivitiesParams) WithLimit(limit *string) *AccountActivitiesGetAccountActivitiesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the account activities get account activities params
func (o *AccountActivitiesGetAccountActivitiesParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the account activities get account activities params
func (o *AccountActivitiesGetAccountActivitiesParams) WithOffset(offset *string) *AccountActivitiesGetAccountActivitiesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the account activities get account activities params
func (o *AccountActivitiesGetAccountActivitiesParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *AccountActivitiesGetAccountActivitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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