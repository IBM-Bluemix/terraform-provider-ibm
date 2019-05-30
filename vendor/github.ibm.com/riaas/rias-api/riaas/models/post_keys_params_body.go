// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostKeysParamsBody KeyTemplate
// swagger:model postKeysParamsBody
type PostKeysParamsBody struct {

	// The user-defined name for this key
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// A public SSH key to be imported into the system
	PublicKey string `json:"public_key,omitempty"`

	// resource group
	ResourceGroup *PostKeysParamsBodyResourceGroup `json:"resource_group,omitempty"`

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`

	// The cryptosystem used by this key
	// Enum: [rsa]
	Type *string `json:"type,omitempty"`
}

// Validate validates this post keys params body
func (m *PostKeysParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostKeysParamsBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PostKeysParamsBody) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if m.ResourceGroup != nil {
		if err := m.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

var postKeysParamsBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rsa"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postKeysParamsBodyTypeTypePropEnum = append(postKeysParamsBodyTypeTypePropEnum, v)
	}
}

const (

	// PostKeysParamsBodyTypeRsa captures enum value "rsa"
	PostKeysParamsBodyTypeRsa string = "rsa"
)

// prop value enum
func (m *PostKeysParamsBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postKeysParamsBodyTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostKeysParamsBody) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostKeysParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostKeysParamsBody) UnmarshalBinary(b []byte) error {
	var res PostKeysParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
