// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkPort network port
// swagger:model NetworkPort
type NetworkPort struct {

	// The escription of the port (not unique or indexable)
	// Required: true
	Description *string `json:"description"`

	// Link to port resource
	Href string `json:"href,omitempty"`

	// The ip address of this port
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// The mac address of the network interface
	// Required: true
	MacAddress *string `json:"macAddress"`

	// The unique Port ID
	// Required: true
	PortID *string `json:"portID"`

	// pvm instance
	PvmInstance *NetworkPortPvmInstance `json:"pvmInstance,omitempty"`

	// Te
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this network port
func (m *NetworkPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMacAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePvmInstance(formats); err != nil {
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

func (m *NetworkPort) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *NetworkPort) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *NetworkPort) validateMacAddress(formats strfmt.Registry) error {

	if err := validate.Required("macAddress", "body", m.MacAddress); err != nil {
		return err
	}

	return nil
}

func (m *NetworkPort) validatePortID(formats strfmt.Registry) error {

	if err := validate.Required("portID", "body", m.PortID); err != nil {
		return err
	}

	return nil
}

func (m *NetworkPort) validatePvmInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.PvmInstance) { // not required
		return nil
	}

	if m.PvmInstance != nil {
		if err := m.PvmInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pvmInstance")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkPort) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkPort) UnmarshalBinary(b []byte) error {
	var res NetworkPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkPortPvmInstance The attached pvm-instance to this port
// swagger:model NetworkPortPvmInstance
type NetworkPortPvmInstance struct {

	// Link to pvm-instance resource
	Href string `json:"href,omitempty"`

	// The attahed pvm-instance ID
	PvmInstanceID string `json:"pvmInstanceID,omitempty"`
}

// Validate validates this network port pvm instance
func (m *NetworkPortPvmInstance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkPortPvmInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkPortPvmInstance) UnmarshalBinary(b []byte) error {
	var res NetworkPortPvmInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
