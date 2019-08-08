// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/power-go-client/power/models"
)

// PcloudTenantsPutReader is a Reader for the PcloudTenantsPut structure.
type PcloudTenantsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTenantsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudTenantsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudTenantsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPcloudTenantsPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudTenantsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudTenantsPutOK creates a PcloudTenantsPutOK with default headers values
func NewPcloudTenantsPutOK() *PcloudTenantsPutOK {
	return &PcloudTenantsPutOK{}
}

/*PcloudTenantsPutOK handles this case with default header values.

OK
*/
type PcloudTenantsPutOK struct {
	Payload *models.Tenant
}

func (o *PcloudTenantsPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudTenantsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsPutBadRequest creates a PcloudTenantsPutBadRequest with default headers values
func NewPcloudTenantsPutBadRequest() *PcloudTenantsPutBadRequest {
	return &PcloudTenantsPutBadRequest{}
}

/*PcloudTenantsPutBadRequest handles this case with default header values.

Bad Request
*/
type PcloudTenantsPutBadRequest struct {
	Payload *models.Error
}

func (o *PcloudTenantsPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTenantsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsPutUnprocessableEntity creates a PcloudTenantsPutUnprocessableEntity with default headers values
func NewPcloudTenantsPutUnprocessableEntity() *PcloudTenantsPutUnprocessableEntity {
	return &PcloudTenantsPutUnprocessableEntity{}
}

/*PcloudTenantsPutUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PcloudTenantsPutUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudTenantsPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudTenantsPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsPutInternalServerError creates a PcloudTenantsPutInternalServerError with default headers values
func NewPcloudTenantsPutInternalServerError() *PcloudTenantsPutInternalServerError {
	return &PcloudTenantsPutInternalServerError{}
}

/*PcloudTenantsPutInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudTenantsPutInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudTenantsPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTenantsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
