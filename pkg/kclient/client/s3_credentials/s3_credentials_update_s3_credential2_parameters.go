// Code generated by go-swagger; DO NOT EDIT.

package s3_credentials

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

// NewS3CredentialsUpdateS3Credential2Params creates a new S3CredentialsUpdateS3Credential2Params object
// with the default values initialized.
func NewS3CredentialsUpdateS3Credential2Params() *S3CredentialsUpdateS3Credential2Params {
	var ()
	return &S3CredentialsUpdateS3Credential2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewS3CredentialsUpdateS3Credential2ParamsWithTimeout creates a new S3CredentialsUpdateS3Credential2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewS3CredentialsUpdateS3Credential2ParamsWithTimeout(timeout time.Duration) *S3CredentialsUpdateS3Credential2Params {
	var ()
	return &S3CredentialsUpdateS3Credential2Params{

		timeout: timeout,
	}
}

// NewS3CredentialsUpdateS3Credential2ParamsWithContext creates a new S3CredentialsUpdateS3Credential2Params object
// with the default values initialized, and the ability to set a context for a request
func NewS3CredentialsUpdateS3Credential2ParamsWithContext(ctx context.Context) *S3CredentialsUpdateS3Credential2Params {
	var ()
	return &S3CredentialsUpdateS3Credential2Params{

		Context: ctx,
	}
}

// NewS3CredentialsUpdateS3Credential2ParamsWithHTTPClient creates a new S3CredentialsUpdateS3Credential2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewS3CredentialsUpdateS3Credential2ParamsWithHTTPClient(client *http.Client) *S3CredentialsUpdateS3Credential2Params {
	var ()
	return &S3CredentialsUpdateS3Credential2Params{
		HTTPClient: client,
	}
}

/*S3CredentialsUpdateS3Credential2Params contains all the parameters to send to the API endpoint
for the s3 credentials update s3 credential2 operation typically these are written to a http.Request
*/
type S3CredentialsUpdateS3Credential2Params struct {

	/*Body*/
	Body *models.AccountS3Credential
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the s3 credentials update s3 credential2 params
func (o *S3CredentialsUpdateS3Credential2Params) WithTimeout(timeout time.Duration) *S3CredentialsUpdateS3Credential2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 credentials update s3 credential2 params
func (o *S3CredentialsUpdateS3Credential2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 credentials update s3 credential2 params
func (o *S3CredentialsUpdateS3Credential2Params) WithContext(ctx context.Context) *S3CredentialsUpdateS3Credential2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 credentials update s3 credential2 params
func (o *S3CredentialsUpdateS3Credential2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 credentials update s3 credential2 params
func (o *S3CredentialsUpdateS3Credential2Params) WithHTTPClient(client *http.Client) *S3CredentialsUpdateS3Credential2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 credentials update s3 credential2 params
func (o *S3CredentialsUpdateS3Credential2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the s3 credentials update s3 credential2 params
func (o *S3CredentialsUpdateS3Credential2Params) WithBody(body *models.AccountS3Credential) *S3CredentialsUpdateS3Credential2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the s3 credentials update s3 credential2 params
func (o *S3CredentialsUpdateS3Credential2Params) SetBody(body *models.AccountS3Credential) {
	o.Body = body
}

// WithID adds the id to the s3 credentials update s3 credential2 params
func (o *S3CredentialsUpdateS3Credential2Params) WithID(id string) *S3CredentialsUpdateS3Credential2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the s3 credentials update s3 credential2 params
func (o *S3CredentialsUpdateS3Credential2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *S3CredentialsUpdateS3Credential2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
