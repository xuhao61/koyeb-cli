// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// ProfileUpdateUserReader is a Reader for the ProfileUpdateUser structure.
type ProfileUpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfileUpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProfileUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProfileUpdateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProfileUpdateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProfileUpdateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProfileUpdateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProfileUpdateUserOK creates a ProfileUpdateUserOK with default headers values
func NewProfileUpdateUserOK() *ProfileUpdateUserOK {
	return &ProfileUpdateUserOK{}
}

/*ProfileUpdateUserOK handles this case with default header values.

A successful response.
*/
type ProfileUpdateUserOK struct {
	Payload *models.AccountUserReply
}

func (o *ProfileUpdateUserOK) Error() string {
	return fmt.Sprintf("[PUT /v1/account/profile][%d] profileUpdateUserOK  %+v", 200, o.Payload)
}

func (o *ProfileUpdateUserOK) GetPayload() *models.AccountUserReply {
	return o.Payload
}

func (o *ProfileUpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountUserReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfileUpdateUserBadRequest creates a ProfileUpdateUserBadRequest with default headers values
func NewProfileUpdateUserBadRequest() *ProfileUpdateUserBadRequest {
	return &ProfileUpdateUserBadRequest{}
}

/*ProfileUpdateUserBadRequest handles this case with default header values.

Validation error
*/
type ProfileUpdateUserBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *ProfileUpdateUserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/account/profile][%d] profileUpdateUserBadRequest  %+v", 400, o.Payload)
}

func (o *ProfileUpdateUserBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *ProfileUpdateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfileUpdateUserForbidden creates a ProfileUpdateUserForbidden with default headers values
func NewProfileUpdateUserForbidden() *ProfileUpdateUserForbidden {
	return &ProfileUpdateUserForbidden{}
}

/*ProfileUpdateUserForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type ProfileUpdateUserForbidden struct {
	Payload *models.CommonError
}

func (o *ProfileUpdateUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/account/profile][%d] profileUpdateUserForbidden  %+v", 403, o.Payload)
}

func (o *ProfileUpdateUserForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ProfileUpdateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfileUpdateUserNotFound creates a ProfileUpdateUserNotFound with default headers values
func NewProfileUpdateUserNotFound() *ProfileUpdateUserNotFound {
	return &ProfileUpdateUserNotFound{}
}

/*ProfileUpdateUserNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type ProfileUpdateUserNotFound struct {
	Payload *models.CommonError
}

func (o *ProfileUpdateUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/account/profile][%d] profileUpdateUserNotFound  %+v", 404, o.Payload)
}

func (o *ProfileUpdateUserNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ProfileUpdateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfileUpdateUserDefault creates a ProfileUpdateUserDefault with default headers values
func NewProfileUpdateUserDefault(code int) *ProfileUpdateUserDefault {
	return &ProfileUpdateUserDefault{
		_statusCode: code,
	}
}

/*ProfileUpdateUserDefault handles this case with default header values.

An unexpected error response
*/
type ProfileUpdateUserDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the profile update user default response
func (o *ProfileUpdateUserDefault) Code() int {
	return o._statusCode
}

func (o *ProfileUpdateUserDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/account/profile][%d] profile_UpdateUser default  %+v", o._statusCode, o.Payload)
}

func (o *ProfileUpdateUserDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *ProfileUpdateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
