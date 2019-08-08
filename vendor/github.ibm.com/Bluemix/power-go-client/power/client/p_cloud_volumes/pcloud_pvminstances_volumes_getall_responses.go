// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/power-go-client/power/models"
)

// PcloudPvminstancesVolumesGetallReader is a Reader for the PcloudPvminstancesVolumesGetall structure.
type PcloudPvminstancesVolumesGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesVolumesGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudPvminstancesVolumesGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudPvminstancesVolumesGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudPvminstancesVolumesGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudPvminstancesVolumesGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudPvminstancesVolumesGetallOK creates a PcloudPvminstancesVolumesGetallOK with default headers values
func NewPcloudPvminstancesVolumesGetallOK() *PcloudPvminstancesVolumesGetallOK {
	return &PcloudPvminstancesVolumesGetallOK{}
}

/*PcloudPvminstancesVolumesGetallOK handles this case with default header values.

OK
*/
type PcloudPvminstancesVolumesGetallOK struct {
	Payload *models.Volumes
}

func (o *PcloudPvminstancesVolumesGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volumes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetallBadRequest creates a PcloudPvminstancesVolumesGetallBadRequest with default headers values
func NewPcloudPvminstancesVolumesGetallBadRequest() *PcloudPvminstancesVolumesGetallBadRequest {
	return &PcloudPvminstancesVolumesGetallBadRequest{}
}

/*PcloudPvminstancesVolumesGetallBadRequest handles this case with default header values.

Bad Request
*/
type PcloudPvminstancesVolumesGetallBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesVolumesGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetallNotFound creates a PcloudPvminstancesVolumesGetallNotFound with default headers values
func NewPcloudPvminstancesVolumesGetallNotFound() *PcloudPvminstancesVolumesGetallNotFound {
	return &PcloudPvminstancesVolumesGetallNotFound{}
}

/*PcloudPvminstancesVolumesGetallNotFound handles this case with default header values.

Not Found
*/
type PcloudPvminstancesVolumesGetallNotFound struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesVolumesGetallNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetallInternalServerError creates a PcloudPvminstancesVolumesGetallInternalServerError with default headers values
func NewPcloudPvminstancesVolumesGetallInternalServerError() *PcloudPvminstancesVolumesGetallInternalServerError {
	return &PcloudPvminstancesVolumesGetallInternalServerError{}
}

/*PcloudPvminstancesVolumesGetallInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudPvminstancesVolumesGetallInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesVolumesGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
