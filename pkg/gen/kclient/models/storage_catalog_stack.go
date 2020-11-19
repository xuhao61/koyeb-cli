// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageCatalogStack A stack entry in the catalog
//
// swagger:model storageCatalogStack
type StorageCatalogStack struct {

	// A longer description of the stack akin to a README in markdown
	Description string `json:"description,omitempty"`

	// A display name for a stack
	DisplayName string `json:"display_name,omitempty"`

	// An optional url to an icon for this stack (What file format?!)
	Icon string `json:"icon,omitempty"`

	// Whether this stack is only a yaml or not (in which case it can be created as a non git managed stack).
	IsYamlOnly bool `json:"is_yaml_only,omitempty"`

	// The name of the tag in a slug valid format
	Name string `json:"name,omitempty"`

	// An short description of the stack (<140 characters)
	ShortDescription string `json:"short_description,omitempty"`

	// A reference to where the stack is defined
	SourceControlRef *StorageSCMReference `json:"source_control_ref,omitempty"`

	// The current status on the stack catalog
	Status StorageBaseCatalogStatus `json:"status,omitempty"`

	// A set of tags to identify the catalog
	Tags []string `json:"tags"`

	// The yaml that is represented the stack.
	Yaml string `json:"yaml,omitempty"`
}

// Validate validates this storage catalog stack
func (m *StorageCatalogStack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceControlRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageCatalogStack) validateSourceControlRef(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceControlRef) { // not required
		return nil
	}

	if m.SourceControlRef != nil {
		if err := m.SourceControlRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_control_ref")
			}
			return err
		}
	}

	return nil
}

func (m *StorageCatalogStack) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageCatalogStack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageCatalogStack) UnmarshalBinary(b []byte) error {
	var res StorageCatalogStack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}