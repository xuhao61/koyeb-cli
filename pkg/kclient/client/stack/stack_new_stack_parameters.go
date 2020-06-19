// Code generated by go-swagger; DO NOT EDIT.

package stack

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

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// NewStackNewStackParams creates a new StackNewStackParams object
// with the default values initialized.
func NewStackNewStackParams() *StackNewStackParams {
	var ()
	return &StackNewStackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStackNewStackParamsWithTimeout creates a new StackNewStackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStackNewStackParamsWithTimeout(timeout time.Duration) *StackNewStackParams {
	var ()
	return &StackNewStackParams{

		timeout: timeout,
	}
}

// NewStackNewStackParamsWithContext creates a new StackNewStackParams object
// with the default values initialized, and the ability to set a context for a request
func NewStackNewStackParamsWithContext(ctx context.Context) *StackNewStackParams {
	var ()
	return &StackNewStackParams{

		Context: ctx,
	}
}

// NewStackNewStackParamsWithHTTPClient creates a new StackNewStackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStackNewStackParamsWithHTTPClient(client *http.Client) *StackNewStackParams {
	var ()
	return &StackNewStackParams{
		HTTPClient: client,
	}
}

/*StackNewStackParams contains all the parameters to send to the API endpoint
for the stack new stack operation typically these are written to a http.Request
*/
type StackNewStackParams struct {

	/*Body*/
	Body *models.StorageStackUpsert

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stack new stack params
func (o *StackNewStackParams) WithTimeout(timeout time.Duration) *StackNewStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stack new stack params
func (o *StackNewStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stack new stack params
func (o *StackNewStackParams) WithContext(ctx context.Context) *StackNewStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stack new stack params
func (o *StackNewStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stack new stack params
func (o *StackNewStackParams) WithHTTPClient(client *http.Client) *StackNewStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stack new stack params
func (o *StackNewStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stack new stack params
func (o *StackNewStackParams) WithBody(body *models.StorageStackUpsert) *StackNewStackParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stack new stack params
func (o *StackNewStackParams) SetBody(body *models.StorageStackUpsert) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StackNewStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
