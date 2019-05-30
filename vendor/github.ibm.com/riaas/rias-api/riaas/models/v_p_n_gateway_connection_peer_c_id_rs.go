// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VPNGatewayConnectionPeerCIDRs v p n gateway connection peer c ID rs
// swagger:model VPNGatewayConnectionPeerCIDRs
type VPNGatewayConnectionPeerCIDRs struct {

	// A collection of peer CIDRs for this resource
	PeerCidrs []string `json:"peer_cidrs"`
}

// Validate validates this v p n gateway connection peer c ID rs
func (m *VPNGatewayConnectionPeerCIDRs) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VPNGatewayConnectionPeerCIDRs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VPNGatewayConnectionPeerCIDRs) UnmarshalBinary(b []byte) error {
	var res VPNGatewayConnectionPeerCIDRs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
