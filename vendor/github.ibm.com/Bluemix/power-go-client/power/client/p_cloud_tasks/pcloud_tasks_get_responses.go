// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/power-go-client/power/models"
)

// PcloudTasksGetReader is a Reader for the PcloudTasksGet structure.
type PcloudTasksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTasksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudTasksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudTasksGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudTasksGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudTasksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudTasksGetOK creates a PcloudTasksGetOK with default headers values
func NewPcloudTasksGetOK() *PcloudTasksGetOK {
	return &PcloudTasksGetOK{}
}

/*PcloudTasksGetOK handles this case with default header values.

OK
*/
type PcloudTasksGetOK struct {
	Payload *models.Task
}

func (o *PcloudTasksGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetOK  %+v", 200, o.Payload)
}

func (o *PcloudTasksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTasksGetBadRequest creates a PcloudTasksGetBadRequest with default headers values
func NewPcloudTasksGetBadRequest() *PcloudTasksGetBadRequest {
	return &PcloudTasksGetBadRequest{}
}

/*PcloudTasksGetBadRequest handles this case with default header values.

Bad Request
*/
type PcloudTasksGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudTasksGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTasksGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTasksGetNotFound creates a PcloudTasksGetNotFound with default headers values
func NewPcloudTasksGetNotFound() *PcloudTasksGetNotFound {
	return &PcloudTasksGetNotFound{}
}

/*PcloudTasksGetNotFound handles this case with default header values.

Not Found
*/
type PcloudTasksGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudTasksGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudTasksGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTasksGetInternalServerError creates a PcloudTasksGetInternalServerError with default headers values
func NewPcloudTasksGetInternalServerError() *PcloudTasksGetInternalServerError {
	return &PcloudTasksGetInternalServerError{}
}

/*PcloudTasksGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudTasksGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudTasksGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTasksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
