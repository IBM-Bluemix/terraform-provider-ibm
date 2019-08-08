// Code generated by go-swagger; DO NOT EDIT.

package open_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/power-go-client/power/models"
)

// ServiceBrokerOpenstacksPostReader is a Reader for the ServiceBrokerOpenstacksPost structure.
type ServiceBrokerOpenstacksPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerOpenstacksPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceBrokerOpenstacksPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewServiceBrokerOpenstacksPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServiceBrokerOpenstacksPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewServiceBrokerOpenstacksPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewServiceBrokerOpenstacksPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewServiceBrokerOpenstacksPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceBrokerOpenstacksPostOK creates a ServiceBrokerOpenstacksPostOK with default headers values
func NewServiceBrokerOpenstacksPostOK() *ServiceBrokerOpenstacksPostOK {
	return &ServiceBrokerOpenstacksPostOK{}
}

/*ServiceBrokerOpenstacksPostOK handles this case with default header values.

OK
*/
type ServiceBrokerOpenstacksPostOK struct {
	Payload *models.OpenStack
}

func (o *ServiceBrokerOpenstacksPostOK) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostCreated creates a ServiceBrokerOpenstacksPostCreated with default headers values
func NewServiceBrokerOpenstacksPostCreated() *ServiceBrokerOpenstacksPostCreated {
	return &ServiceBrokerOpenstacksPostCreated{}
}

/*ServiceBrokerOpenstacksPostCreated handles this case with default header values.

Created
*/
type ServiceBrokerOpenstacksPostCreated struct {
	Payload *models.OpenStack
}

func (o *ServiceBrokerOpenstacksPostCreated) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostCreated  %+v", 201, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostBadRequest creates a ServiceBrokerOpenstacksPostBadRequest with default headers values
func NewServiceBrokerOpenstacksPostBadRequest() *ServiceBrokerOpenstacksPostBadRequest {
	return &ServiceBrokerOpenstacksPostBadRequest{}
}

/*ServiceBrokerOpenstacksPostBadRequest handles this case with default header values.

Bad Request
*/
type ServiceBrokerOpenstacksPostBadRequest struct {
	Payload *models.Error
}

func (o *ServiceBrokerOpenstacksPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostConflict creates a ServiceBrokerOpenstacksPostConflict with default headers values
func NewServiceBrokerOpenstacksPostConflict() *ServiceBrokerOpenstacksPostConflict {
	return &ServiceBrokerOpenstacksPostConflict{}
}

/*ServiceBrokerOpenstacksPostConflict handles this case with default header values.

Conflict
*/
type ServiceBrokerOpenstacksPostConflict struct {
	Payload *models.Error
}

func (o *ServiceBrokerOpenstacksPostConflict) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostConflict  %+v", 409, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostUnprocessableEntity creates a ServiceBrokerOpenstacksPostUnprocessableEntity with default headers values
func NewServiceBrokerOpenstacksPostUnprocessableEntity() *ServiceBrokerOpenstacksPostUnprocessableEntity {
	return &ServiceBrokerOpenstacksPostUnprocessableEntity{}
}

/*ServiceBrokerOpenstacksPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type ServiceBrokerOpenstacksPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostInternalServerError creates a ServiceBrokerOpenstacksPostInternalServerError with default headers values
func NewServiceBrokerOpenstacksPostInternalServerError() *ServiceBrokerOpenstacksPostInternalServerError {
	return &ServiceBrokerOpenstacksPostInternalServerError{}
}

/*ServiceBrokerOpenstacksPostInternalServerError handles this case with default header values.

Internal Server Error
*/
type ServiceBrokerOpenstacksPostInternalServerError struct {
	Payload *models.Error
}

func (o *ServiceBrokerOpenstacksPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
