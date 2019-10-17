// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PatchVolumesVolumeIDSnapshotsIDReader is a Reader for the PatchVolumesVolumeIDSnapshotsID structure.
type PatchVolumesVolumeIDSnapshotsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchVolumesVolumeIDSnapshotsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchVolumesVolumeIDSnapshotsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchVolumesVolumeIDSnapshotsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchVolumesVolumeIDSnapshotsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPatchVolumesVolumeIDSnapshotsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchVolumesVolumeIDSnapshotsIDOK creates a PatchVolumesVolumeIDSnapshotsIDOK with default headers values
func NewPatchVolumesVolumeIDSnapshotsIDOK() *PatchVolumesVolumeIDSnapshotsIDOK {
	return &PatchVolumesVolumeIDSnapshotsIDOK{}
}

/*PatchVolumesVolumeIDSnapshotsIDOK handles this case with default header values.

dummy
*/
type PatchVolumesVolumeIDSnapshotsIDOK struct {
	Payload *models.VolumeSnapshot
}

func (o *PatchVolumesVolumeIDSnapshotsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{volume_id}/snapshots/{id}][%d] patchVolumesVolumeIdSnapshotsIdOK  %+v", 200, o.Payload)
}

func (o *PatchVolumesVolumeIDSnapshotsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVolumesVolumeIDSnapshotsIDBadRequest creates a PatchVolumesVolumeIDSnapshotsIDBadRequest with default headers values
func NewPatchVolumesVolumeIDSnapshotsIDBadRequest() *PatchVolumesVolumeIDSnapshotsIDBadRequest {
	return &PatchVolumesVolumeIDSnapshotsIDBadRequest{}
}

/*PatchVolumesVolumeIDSnapshotsIDBadRequest handles this case with default header values.

error
*/
type PatchVolumesVolumeIDSnapshotsIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchVolumesVolumeIDSnapshotsIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{volume_id}/snapshots/{id}][%d] patchVolumesVolumeIdSnapshotsIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchVolumesVolumeIDSnapshotsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVolumesVolumeIDSnapshotsIDNotFound creates a PatchVolumesVolumeIDSnapshotsIDNotFound with default headers values
func NewPatchVolumesVolumeIDSnapshotsIDNotFound() *PatchVolumesVolumeIDSnapshotsIDNotFound {
	return &PatchVolumesVolumeIDSnapshotsIDNotFound{}
}

/*PatchVolumesVolumeIDSnapshotsIDNotFound handles this case with default header values.

error
*/
type PatchVolumesVolumeIDSnapshotsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *PatchVolumesVolumeIDSnapshotsIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{volume_id}/snapshots/{id}][%d] patchVolumesVolumeIdSnapshotsIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchVolumesVolumeIDSnapshotsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVolumesVolumeIDSnapshotsIDInternalServerError creates a PatchVolumesVolumeIDSnapshotsIDInternalServerError with default headers values
func NewPatchVolumesVolumeIDSnapshotsIDInternalServerError() *PatchVolumesVolumeIDSnapshotsIDInternalServerError {
	return &PatchVolumesVolumeIDSnapshotsIDInternalServerError{}
}

/*PatchVolumesVolumeIDSnapshotsIDInternalServerError handles this case with default header values.

error
*/
type PatchVolumesVolumeIDSnapshotsIDInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *PatchVolumesVolumeIDSnapshotsIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{volume_id}/snapshots/{id}][%d] patchVolumesVolumeIdSnapshotsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchVolumesVolumeIDSnapshotsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
