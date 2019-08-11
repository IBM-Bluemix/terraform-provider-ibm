// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/power-go-client/power/models"
)

// PcloudPvminstancesNetworksGetReader is a Reader for the PcloudPvminstancesNetworksGet structure.
type PcloudPvminstancesNetworksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesNetworksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudPvminstancesNetworksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPcloudPvminstancesNetworksGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudPvminstancesNetworksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudPvminstancesNetworksGetOK creates a PcloudPvminstancesNetworksGetOK with default headers values
func NewPcloudPvminstancesNetworksGetOK() *PcloudPvminstancesNetworksGetOK {
	return &PcloudPvminstancesNetworksGetOK{}
}

/*PcloudPvminstancesNetworksGetOK handles this case with default header values.

OK
*/
type PcloudPvminstancesNetworksGetOK struct {
	Payload *models.PVMInstanceNetwork
}

func (o *PcloudPvminstancesNetworksGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstanceNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetNotFound creates a PcloudPvminstancesNetworksGetNotFound with default headers values
func NewPcloudPvminstancesNetworksGetNotFound() *PcloudPvminstancesNetworksGetNotFound {
	return &PcloudPvminstancesNetworksGetNotFound{}
}

/*PcloudPvminstancesNetworksGetNotFound handles this case with default header values.

Not Found
*/
type PcloudPvminstancesNetworksGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesNetworksGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetInternalServerError creates a PcloudPvminstancesNetworksGetInternalServerError with default headers values
func NewPcloudPvminstancesNetworksGetInternalServerError() *PcloudPvminstancesNetworksGetInternalServerError {
	return &PcloudPvminstancesNetworksGetInternalServerError{}
}

/*PcloudPvminstancesNetworksGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudPvminstancesNetworksGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesNetworksGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
