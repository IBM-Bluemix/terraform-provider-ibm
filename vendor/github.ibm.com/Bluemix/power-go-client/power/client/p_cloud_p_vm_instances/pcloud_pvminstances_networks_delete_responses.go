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

// PcloudPvminstancesNetworksDeleteReader is a Reader for the PcloudPvminstancesNetworksDelete structure.
type PcloudPvminstancesNetworksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesNetworksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudPvminstancesNetworksDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudPvminstancesNetworksDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 410:
		result := NewPcloudPvminstancesNetworksDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudPvminstancesNetworksDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudPvminstancesNetworksDeleteOK creates a PcloudPvminstancesNetworksDeleteOK with default headers values
func NewPcloudPvminstancesNetworksDeleteOK() *PcloudPvminstancesNetworksDeleteOK {
	return &PcloudPvminstancesNetworksDeleteOK{}
}

/*PcloudPvminstancesNetworksDeleteOK handles this case with default header values.

OK
*/
type PcloudPvminstancesNetworksDeleteOK struct {
	Payload models.Object
}

func (o *PcloudPvminstancesNetworksDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesNetworksDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksDeleteBadRequest creates a PcloudPvminstancesNetworksDeleteBadRequest with default headers values
func NewPcloudPvminstancesNetworksDeleteBadRequest() *PcloudPvminstancesNetworksDeleteBadRequest {
	return &PcloudPvminstancesNetworksDeleteBadRequest{}
}

/*PcloudPvminstancesNetworksDeleteBadRequest handles this case with default header values.

Bad Request
*/
type PcloudPvminstancesNetworksDeleteBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesNetworksDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesNetworksDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksDeleteGone creates a PcloudPvminstancesNetworksDeleteGone with default headers values
func NewPcloudPvminstancesNetworksDeleteGone() *PcloudPvminstancesNetworksDeleteGone {
	return &PcloudPvminstancesNetworksDeleteGone{}
}

/*PcloudPvminstancesNetworksDeleteGone handles this case with default header values.

Gone
*/
type PcloudPvminstancesNetworksDeleteGone struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesNetworksDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudPvminstancesNetworksDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksDeleteInternalServerError creates a PcloudPvminstancesNetworksDeleteInternalServerError with default headers values
func NewPcloudPvminstancesNetworksDeleteInternalServerError() *PcloudPvminstancesNetworksDeleteInternalServerError {
	return &PcloudPvminstancesNetworksDeleteInternalServerError{}
}

/*PcloudPvminstancesNetworksDeleteInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudPvminstancesNetworksDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesNetworksDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesNetworksDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
