// Code generated by go-swagger; DO NOT EDIT.

package hardware_platforms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	"github.com/go-openapi/strfmt"
)

// New creates a new hardware platforms API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hardware platforms API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ServiceBrokerHardwareplatformsGet availables hardware statistics and limits
*/
func (a *Client) ServiceBrokerHardwareplatformsGet(params *ServiceBrokerHardwareplatformsGetParams) (*ServiceBrokerHardwareplatformsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerHardwareplatformsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.hardwareplatforms.get",
		Method:             "GET",
		PathPattern:        "/broker/v1/hardware-platforms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerHardwareplatformsGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerHardwareplatformsGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
