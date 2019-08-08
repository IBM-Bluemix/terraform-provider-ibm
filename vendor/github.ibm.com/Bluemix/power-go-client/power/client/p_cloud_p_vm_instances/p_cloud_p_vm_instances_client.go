// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"log"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new p cloud p vm instances API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p cloud p vm instances API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PcloudPvminstancesActionPost performs an action start stop reboot on a p VM instance
*/
func (a *Client) PcloudPvminstancesActionPost(params *PcloudPvminstancesActionPostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesActionPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesActionPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.action.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesActionPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesActionPostOK), nil

}

/*
PcloudPvminstancesCapturePost captures a p VM instance and create a deployable image
*/
func (a *Client) PcloudPvminstancesCapturePost(params *PcloudPvminstancesCapturePostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesCapturePostOK, *PcloudPvminstancesCapturePostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesCapturePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.capture.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesCapturePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudPvminstancesCapturePostOK:
		return value, nil, nil
	case *PcloudPvminstancesCapturePostAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PcloudPvminstancesConsolePost generates the no v n c console URL
*/
func (a *Client) PcloudPvminstancesConsolePost(params *PcloudPvminstancesConsolePostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesConsolePostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesConsolePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.console.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesConsolePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesConsolePostCreated), nil

}

/*
PcloudPvminstancesDelete deletes a p cloud p VM instance
*/
func (a *Client) PcloudPvminstancesDelete(params *PcloudPvminstancesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesDeleteOK), nil

}

/*
PcloudPvminstancesGet gets a p VM instance s current state information
*/
func (a *Client) PcloudPvminstancesGet(params *PcloudPvminstancesGetParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesGetOK), nil

}

/*
PcloudPvminstancesGetall gets all the pvm instances for this cloud instance
*/
func (a *Client) PcloudPvminstancesGetall(params *PcloudPvminstancesGetallParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesGetallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesGetallOK), nil

}

/*
PcloudPvminstancesPost creates a new power VM instance
*/
func (a *Client) PcloudPvminstancesPost(params *PcloudPvminstancesPostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesPostOK, *PcloudPvminstancesPostCreated, *PcloudPvminstancesPostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesPostParams()
	}

	log.Printf("Printing the params that are passed to the create instance calls %+v ",params.Body)
	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudPvminstancesPostOK:
		return value, nil, nil, nil
	case *PcloudPvminstancesPostCreated:
		return nil, value, nil, nil
	case *PcloudPvminstancesPostAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
PcloudPvminstancesPut updates a p cloud p VM instance
*/
func (a *Client) PcloudPvminstancesPut(params *PcloudPvminstancesPutParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesPutAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesPutParams()
	}
	log.Printf("Calling this update with the following params %+v ",params)
	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesPutAccepted), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
