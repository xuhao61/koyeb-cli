// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/gen/kclient/models"
)

// ActivityListNotificationsReader is a Reader for the ActivityListNotifications structure.
type ActivityListNotificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActivityListNotificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActivityListNotificationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewActivityListNotificationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewActivityListNotificationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewActivityListNotificationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewActivityListNotificationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActivityListNotificationsOK creates a ActivityListNotificationsOK with default headers values
func NewActivityListNotificationsOK() *ActivityListNotificationsOK {
	return &ActivityListNotificationsOK{}
}

/*ActivityListNotificationsOK handles this case with default header values.

A successful response.
*/
type ActivityListNotificationsOK struct {
	Payload *models.ActivityNotificationList
}

func (o *ActivityListNotificationsOK) Error() string {
	return fmt.Sprintf("[GET /v1/notifications][%d] activityListNotificationsOK  %+v", 200, o.Payload)
}

func (o *ActivityListNotificationsOK) GetPayload() *models.ActivityNotificationList {
	return o.Payload
}

func (o *ActivityListNotificationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActivityNotificationList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivityListNotificationsBadRequest creates a ActivityListNotificationsBadRequest with default headers values
func NewActivityListNotificationsBadRequest() *ActivityListNotificationsBadRequest {
	return &ActivityListNotificationsBadRequest{}
}

/*ActivityListNotificationsBadRequest handles this case with default header values.

Validation error
*/
type ActivityListNotificationsBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *ActivityListNotificationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/notifications][%d] activityListNotificationsBadRequest  %+v", 400, o.Payload)
}

func (o *ActivityListNotificationsBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *ActivityListNotificationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivityListNotificationsForbidden creates a ActivityListNotificationsForbidden with default headers values
func NewActivityListNotificationsForbidden() *ActivityListNotificationsForbidden {
	return &ActivityListNotificationsForbidden{}
}

/*ActivityListNotificationsForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type ActivityListNotificationsForbidden struct {
	Payload *models.CommonError
}

func (o *ActivityListNotificationsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/notifications][%d] activityListNotificationsForbidden  %+v", 403, o.Payload)
}

func (o *ActivityListNotificationsForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ActivityListNotificationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivityListNotificationsNotFound creates a ActivityListNotificationsNotFound with default headers values
func NewActivityListNotificationsNotFound() *ActivityListNotificationsNotFound {
	return &ActivityListNotificationsNotFound{}
}

/*ActivityListNotificationsNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type ActivityListNotificationsNotFound struct {
	Payload *models.CommonError
}

func (o *ActivityListNotificationsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/notifications][%d] activityListNotificationsNotFound  %+v", 404, o.Payload)
}

func (o *ActivityListNotificationsNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ActivityListNotificationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivityListNotificationsDefault creates a ActivityListNotificationsDefault with default headers values
func NewActivityListNotificationsDefault(code int) *ActivityListNotificationsDefault {
	return &ActivityListNotificationsDefault{
		_statusCode: code,
	}
}

/*ActivityListNotificationsDefault handles this case with default header values.

An unexpected error response
*/
type ActivityListNotificationsDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the activity list notifications default response
func (o *ActivityListNotificationsDefault) Code() int {
	return o._statusCode
}

func (o *ActivityListNotificationsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/notifications][%d] activity_ListNotifications default  %+v", o._statusCode, o.Payload)
}

func (o *ActivityListNotificationsDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *ActivityListNotificationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}