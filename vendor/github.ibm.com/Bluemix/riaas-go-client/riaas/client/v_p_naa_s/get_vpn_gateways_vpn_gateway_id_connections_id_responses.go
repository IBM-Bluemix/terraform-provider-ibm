// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetVpnGatewaysVpnGatewayIDConnectionsIDReader is a Reader for the GetVpnGatewaysVpnGatewayIDConnectionsID structure.
type GetVpnGatewaysVpnGatewayIDConnectionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVpnGatewaysVpnGatewayIDConnectionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetVpnGatewaysVpnGatewayIDConnectionsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDOK creates a GetVpnGatewaysVpnGatewayIDConnectionsIDOK with default headers values
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDOK() *GetVpnGatewaysVpnGatewayIDConnectionsIDOK {
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDOK{}
}

/*GetVpnGatewaysVpnGatewayIDConnectionsIDOK handles this case with default header values.

The VPN connection was retrieved successfully.
*/
type GetVpnGatewaysVpnGatewayIDConnectionsIDOK struct {
	Payload *models.VPNGatewayConnection
}

func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDOK) Error() string {
	return fmt.Sprintf("[GET /vpn_gateways/{vpn_gateway_id}/connections/{id}][%d] getVpnGatewaysVpnGatewayIdConnectionsIdOK  %+v", 200, o.Payload)
}

func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VPNGatewayConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDNotFound creates a GetVpnGatewaysVpnGatewayIDConnectionsIDNotFound with default headers values
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDNotFound() *GetVpnGatewaysVpnGatewayIDConnectionsIDNotFound {
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDNotFound{}
}

/*GetVpnGatewaysVpnGatewayIDConnectionsIDNotFound handles this case with default header values.

A VPN connection with the specified identifier could not be found.
*/
type GetVpnGatewaysVpnGatewayIDConnectionsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /vpn_gateways/{vpn_gateway_id}/connections/{id}][%d] getVpnGatewaysVpnGatewayIdConnectionsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
