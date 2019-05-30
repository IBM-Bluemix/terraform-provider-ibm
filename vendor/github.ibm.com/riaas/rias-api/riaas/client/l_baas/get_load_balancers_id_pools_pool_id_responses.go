// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetLoadBalancersIDPoolsPoolIDReader is a Reader for the GetLoadBalancersIDPoolsPoolID structure.
type GetLoadBalancersIDPoolsPoolIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoadBalancersIDPoolsPoolIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLoadBalancersIDPoolsPoolIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetLoadBalancersIDPoolsPoolIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLoadBalancersIDPoolsPoolIDOK creates a GetLoadBalancersIDPoolsPoolIDOK with default headers values
func NewGetLoadBalancersIDPoolsPoolIDOK() *GetLoadBalancersIDPoolsPoolIDOK {
	return &GetLoadBalancersIDPoolsPoolIDOK{}
}

/*GetLoadBalancersIDPoolsPoolIDOK handles this case with default header values.

The pool was retrieved successfully.
*/
type GetLoadBalancersIDPoolsPoolIDOK struct {
	Payload *models.Pool
}

func (o *GetLoadBalancersIDPoolsPoolIDOK) Error() string {
	return fmt.Sprintf("[GET /load_balancers/{id}/pools/{pool_id}][%d] getLoadBalancersIdPoolsPoolIdOK  %+v", 200, o.Payload)
}

func (o *GetLoadBalancersIDPoolsPoolIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoadBalancersIDPoolsPoolIDNotFound creates a GetLoadBalancersIDPoolsPoolIDNotFound with default headers values
func NewGetLoadBalancersIDPoolsPoolIDNotFound() *GetLoadBalancersIDPoolsPoolIDNotFound {
	return &GetLoadBalancersIDPoolsPoolIDNotFound{}
}

/*GetLoadBalancersIDPoolsPoolIDNotFound handles this case with default header values.

A load balancer with the specified identifier could not be found.
*/
type GetLoadBalancersIDPoolsPoolIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetLoadBalancersIDPoolsPoolIDNotFound) Error() string {
	return fmt.Sprintf("[GET /load_balancers/{id}/pools/{pool_id}][%d] getLoadBalancersIdPoolsPoolIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLoadBalancersIDPoolsPoolIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
