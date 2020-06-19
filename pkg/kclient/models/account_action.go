// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AccountAction account action
// swagger:model accountAction
type AccountAction string

const (

	// AccountActionSignin captures enum value "signin"
	AccountActionSignin AccountAction = "signin"

	// AccountActionSignup captures enum value "signup"
	AccountActionSignup AccountAction = "signup"

	// AccountActionRegister captures enum value "register"
	AccountActionRegister AccountAction = "register"
)

// for schema
var accountActionEnum []interface{}

func init() {
	var res []AccountAction
	if err := json.Unmarshal([]byte(`["signin","signup","register"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountActionEnum = append(accountActionEnum, v)
	}
}

func (m AccountAction) validateAccountActionEnum(path, location string, value AccountAction) error {
	if err := validate.Enum(path, location, value, accountActionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this account action
func (m AccountAction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccountActionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
