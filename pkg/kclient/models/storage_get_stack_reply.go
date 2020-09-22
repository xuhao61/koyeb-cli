// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageGetStackReply storage get stack reply
// swagger:model storageGetStackReply
type StorageGetStackReply struct {

	// stack
	Stack *StorageStack `json:"stack,omitempty"`
}

// Validate validates this storage get stack reply
func (m *StorageGetStackReply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageGetStackReply) validateStack(formats strfmt.Registry) error {

	if swag.IsZero(m.Stack) { // not required
		return nil
	}

	if m.Stack != nil {
		if err := m.Stack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageGetStackReply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageGetStackReply) UnmarshalBinary(b []byte) error {
	var res StorageGetStackReply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}