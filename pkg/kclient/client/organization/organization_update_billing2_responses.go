// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// OrganizationUpdateBilling2Reader is a Reader for the OrganizationUpdateBilling2 structure.
type OrganizationUpdateBilling2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationUpdateBilling2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationUpdateBilling2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationUpdateBilling2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationUpdateBilling2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationUpdateBilling2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOrganizationUpdateBilling2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationUpdateBilling2OK creates a OrganizationUpdateBilling2OK with default headers values
func NewOrganizationUpdateBilling2OK() *OrganizationUpdateBilling2OK {
	return &OrganizationUpdateBilling2OK{}
}

/*OrganizationUpdateBilling2OK handles this case with default header values.

A successful response.
*/
type OrganizationUpdateBilling2OK struct {
	Payload *models.AccountBillingInfoReply
}

func (o *OrganizationUpdateBilling2OK) Error() string {
	return fmt.Sprintf("[PATCH /v1/account/billing][%d] organizationUpdateBilling2OK  %+v", 200, o.Payload)
}

func (o *OrganizationUpdateBilling2OK) GetPayload() *models.AccountBillingInfoReply {
	return o.Payload
}

func (o *OrganizationUpdateBilling2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountBillingInfoReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationUpdateBilling2BadRequest creates a OrganizationUpdateBilling2BadRequest with default headers values
func NewOrganizationUpdateBilling2BadRequest() *OrganizationUpdateBilling2BadRequest {
	return &OrganizationUpdateBilling2BadRequest{}
}

/*OrganizationUpdateBilling2BadRequest handles this case with default header values.

Validation error
*/
type OrganizationUpdateBilling2BadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *OrganizationUpdateBilling2BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/account/billing][%d] organizationUpdateBilling2BadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationUpdateBilling2BadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *OrganizationUpdateBilling2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationUpdateBilling2Forbidden creates a OrganizationUpdateBilling2Forbidden with default headers values
func NewOrganizationUpdateBilling2Forbidden() *OrganizationUpdateBilling2Forbidden {
	return &OrganizationUpdateBilling2Forbidden{}
}

/*OrganizationUpdateBilling2Forbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type OrganizationUpdateBilling2Forbidden struct {
	Payload *models.CommonError
}

func (o *OrganizationUpdateBilling2Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/account/billing][%d] organizationUpdateBilling2Forbidden  %+v", 403, o.Payload)
}

func (o *OrganizationUpdateBilling2Forbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *OrganizationUpdateBilling2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationUpdateBilling2NotFound creates a OrganizationUpdateBilling2NotFound with default headers values
func NewOrganizationUpdateBilling2NotFound() *OrganizationUpdateBilling2NotFound {
	return &OrganizationUpdateBilling2NotFound{}
}

/*OrganizationUpdateBilling2NotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type OrganizationUpdateBilling2NotFound struct {
	Payload *models.CommonError
}

func (o *OrganizationUpdateBilling2NotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/account/billing][%d] organizationUpdateBilling2NotFound  %+v", 404, o.Payload)
}

func (o *OrganizationUpdateBilling2NotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *OrganizationUpdateBilling2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationUpdateBilling2Default creates a OrganizationUpdateBilling2Default with default headers values
func NewOrganizationUpdateBilling2Default(code int) *OrganizationUpdateBilling2Default {
	return &OrganizationUpdateBilling2Default{
		_statusCode: code,
	}
}

/*OrganizationUpdateBilling2Default handles this case with default header values.

An unexpected error response
*/
type OrganizationUpdateBilling2Default struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the organization update billing2 default response
func (o *OrganizationUpdateBilling2Default) Code() int {
	return o._statusCode
}

func (o *OrganizationUpdateBilling2Default) Error() string {
	return fmt.Sprintf("[PATCH /v1/account/billing][%d] organization_UpdateBilling2 default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationUpdateBilling2Default) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *OrganizationUpdateBilling2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}