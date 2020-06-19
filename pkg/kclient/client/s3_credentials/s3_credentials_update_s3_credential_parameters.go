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

// NewS3CredentialsUpdateS3CredentialParams creates a new S3CredentialsUpdateS3CredentialParams object
// with the default values initialized.
func NewS3CredentialsUpdateS3CredentialParams() *S3CredentialsUpdateS3CredentialParams {
	var ()
	return &S3CredentialsUpdateS3CredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewS3CredentialsUpdateS3CredentialParamsWithTimeout creates a new S3CredentialsUpdateS3CredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewS3CredentialsUpdateS3CredentialParamsWithTimeout(timeout time.Duration) *S3CredentialsUpdateS3CredentialParams {
	var ()
	return &S3CredentialsUpdateS3CredentialParams{

		timeout: timeout,
	}
}

// NewS3CredentialsUpdateS3CredentialParamsWithContext creates a new S3CredentialsUpdateS3CredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewS3CredentialsUpdateS3CredentialParamsWithContext(ctx context.Context) *S3CredentialsUpdateS3CredentialParams {
	var ()
	return &S3CredentialsUpdateS3CredentialParams{

		Context: ctx,
	}
}

// NewS3CredentialsUpdateS3CredentialParamsWithHTTPClient creates a new S3CredentialsUpdateS3CredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewS3CredentialsUpdateS3CredentialParamsWithHTTPClient(client *http.Client) *S3CredentialsUpdateS3CredentialParams {
	var ()
	return &S3CredentialsUpdateS3CredentialParams{
		HTTPClient: client,
	}
}

/*S3CredentialsUpdateS3CredentialParams contains all the parameters to send to the API endpoint
for the s3 credentials update s3 credential operation typically these are written to a http.Request
*/
type S3CredentialsUpdateS3CredentialParams struct {

	/*Body*/
	Body *models.AccountS3Credential
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the s3 credentials update s3 credential params
func (o *S3CredentialsUpdateS3CredentialParams) WithTimeout(timeout time.Duration) *S3CredentialsUpdateS3CredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 credentials update s3 credential params
func (o *S3CredentialsUpdateS3CredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 credentials update s3 credential params
func (o *S3CredentialsUpdateS3CredentialParams) WithContext(ctx context.Context) *S3CredentialsUpdateS3CredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 credentials update s3 credential params
func (o *S3CredentialsUpdateS3CredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 credentials update s3 credential params
func (o *S3CredentialsUpdateS3CredentialParams) WithHTTPClient(client *http.Client) *S3CredentialsUpdateS3CredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 credentials update s3 credential params
func (o *S3CredentialsUpdateS3CredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the s3 credentials update s3 credential params
func (o *S3CredentialsUpdateS3CredentialParams) WithBody(body *models.AccountS3Credential) *S3CredentialsUpdateS3CredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the s3 credentials update s3 credential params
func (o *S3CredentialsUpdateS3CredentialParams) SetBody(body *models.AccountS3Credential) {
	o.Body = body
}

// WithID adds the id to the s3 credentials update s3 credential params
func (o *S3CredentialsUpdateS3CredentialParams) WithID(id string) *S3CredentialsUpdateS3CredentialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the s3 credentials update s3 credential params
func (o *S3CredentialsUpdateS3CredentialParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *S3CredentialsUpdateS3CredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
