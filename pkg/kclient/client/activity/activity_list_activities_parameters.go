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

// NewActivityListActivitiesParams creates a new ActivityListActivitiesParams object
// with the default values initialized.
func NewActivityListActivitiesParams() *ActivityListActivitiesParams {
	var ()
	return &ActivityListActivitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActivityListActivitiesParamsWithTimeout creates a new ActivityListActivitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActivityListActivitiesParamsWithTimeout(timeout time.Duration) *ActivityListActivitiesParams {
	var ()
	return &ActivityListActivitiesParams{

		timeout: timeout,
	}
}

// NewActivityListActivitiesParamsWithContext creates a new ActivityListActivitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewActivityListActivitiesParamsWithContext(ctx context.Context) *ActivityListActivitiesParams {
	var ()
	return &ActivityListActivitiesParams{

		Context: ctx,
	}
}

// NewActivityListActivitiesParamsWithHTTPClient creates a new ActivityListActivitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActivityListActivitiesParamsWithHTTPClient(client *http.Client) *ActivityListActivitiesParams {
	var ()
	return &ActivityListActivitiesParams{
		HTTPClient: client,
	}
}

/*ActivityListActivitiesParams contains all the parameters to send to the API endpoint
for the activity list activities operation typically these are written to a http.Request
*/
type ActivityListActivitiesParams struct {

	/*Limit*/
	Limit *string
	/*Offset*/
	Offset *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the activity list activities params
func (o *ActivityListActivitiesParams) WithTimeout(timeout time.Duration) *ActivityListActivitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the activity list activities params
func (o *ActivityListActivitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the activity list activities params
func (o *ActivityListActivitiesParams) WithContext(ctx context.Context) *ActivityListActivitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the activity list activities params
func (o *ActivityListActivitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the activity list activities params
func (o *ActivityListActivitiesParams) WithHTTPClient(client *http.Client) *ActivityListActivitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the activity list activities params
func (o *ActivityListActivitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the activity list activities params
func (o *ActivityListActivitiesParams) WithLimit(limit *string) *ActivityListActivitiesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the activity list activities params
func (o *ActivityListActivitiesParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the activity list activities params
func (o *ActivityListActivitiesParams) WithOffset(offset *string) *ActivityListActivitiesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the activity list activities params
func (o *ActivityListActivitiesParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ActivityListActivitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
