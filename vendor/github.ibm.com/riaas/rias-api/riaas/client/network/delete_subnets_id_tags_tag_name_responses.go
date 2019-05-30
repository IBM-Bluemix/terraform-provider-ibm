// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// DeleteSubnetsIDTagsTagNameReader is a Reader for the DeleteSubnetsIDTagsTagName structure.
type DeleteSubnetsIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSubnetsIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSubnetsIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteSubnetsIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteSubnetsIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewDeleteSubnetsIDTagsTagNameNoContent creates a DeleteSubnetsIDTagsTagNameNoContent with default headers values
func NewDeleteSubnetsIDTagsTagNameNoContent() *DeleteSubnetsIDTagsTagNameNoContent {
	return &DeleteSubnetsIDTagsTagNameNoContent{}
}

/*DeleteSubnetsIDTagsTagNameNoContent handles this case with default header values.

error
*/
type DeleteSubnetsIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *DeleteSubnetsIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /subnets/{id}/tags/{tag_name}][%d] deleteSubnetsIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *DeleteSubnetsIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSubnetsIDTagsTagNameNotFound creates a DeleteSubnetsIDTagsTagNameNotFound with default headers values
func NewDeleteSubnetsIDTagsTagNameNotFound() *DeleteSubnetsIDTagsTagNameNotFound {
	return &DeleteSubnetsIDTagsTagNameNotFound{}
}

/*DeleteSubnetsIDTagsTagNameNotFound handles this case with default header values.

error
*/
type DeleteSubnetsIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteSubnetsIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /subnets/{id}/tags/{tag_name}][%d] deleteSubnetsIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSubnetsIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSubnetsIDTagsTagNameDefault creates a DeleteSubnetsIDTagsTagNameDefault with default headers values
func NewDeleteSubnetsIDTagsTagNameDefault(code int) *DeleteSubnetsIDTagsTagNameDefault {
	return &DeleteSubnetsIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*DeleteSubnetsIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type DeleteSubnetsIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the delete subnets ID tags tag name default response
func (o *DeleteSubnetsIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSubnetsIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /subnets/{id}/tags/{tag_name}][%d] DeleteSubnetsIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSubnetsIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
