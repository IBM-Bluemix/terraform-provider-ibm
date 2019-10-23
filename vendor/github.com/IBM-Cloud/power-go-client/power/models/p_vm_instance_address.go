// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PVMInstanceAddress deprecated - replaced by PVMInstanceNetwork
// swagger:model PVMInstanceAddress
type PVMInstanceAddress struct {
	PVMInstanceNetwork
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PVMInstanceAddress) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PVMInstanceNetwork
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PVMInstanceNetwork = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PVMInstanceAddress) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.PVMInstanceNetwork)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this p VM instance address
func (m *PVMInstanceAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceAddress) UnmarshalBinary(b []byte) error {
	var res PVMInstanceAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
