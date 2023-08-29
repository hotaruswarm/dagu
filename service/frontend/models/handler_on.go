// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HandlerOn handler on
//
// swagger:model handlerOn
type HandlerOn struct {

	// cancel
	Cancel string `json:"Cancel,omitempty"`

	// exit
	Exit string `json:"Exit,omitempty"`

	// failure
	Failure string `json:"Failure,omitempty"`

	// success
	Success string `json:"Success,omitempty"`
}

// Validate validates this handler on
func (m *HandlerOn) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this handler on based on context it is used
func (m *HandlerOn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HandlerOn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HandlerOn) UnmarshalBinary(b []byte) error {
	var res HandlerOn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}