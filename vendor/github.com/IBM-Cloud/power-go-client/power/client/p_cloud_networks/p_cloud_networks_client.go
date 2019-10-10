// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	"github.com/go-openapi/strfmt"
)

// New creates a new p cloud networks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p cloud networks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PcloudNetworksDelete deletes a network
*/
func (a *Client) PcloudNetworksDelete(params *PcloudNetworksDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksDeleteOK), nil

}

/*
PcloudNetworksGet gets a network s current state information
*/
func (a *Client) PcloudNetworksGet(params *PcloudNetworksGetParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksGetOK), nil

}

/*
PcloudNetworksGetall gets all networks in this cloud instance
*/
func (a *Client) PcloudNetworksGetall(params *PcloudNetworksGetallParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksGetallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksGetallOK), nil

}

/*
PcloudNetworksPost creates a new network
*/
func (a *Client) PcloudNetworksPost(params *PcloudNetworksPostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksPostOK, *PcloudNetworksPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudNetworksPostOK:
		return value, nil, nil
	case *PcloudNetworksPostCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PcloudNetworksPut updates a network
*/
func (a *Client) PcloudNetworksPut(params *PcloudNetworksPutParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksPutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
