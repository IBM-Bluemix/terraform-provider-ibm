// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetSecurityGroupsSecurityGroupIDRulesIDReader is a Reader for the GetSecurityGroupsSecurityGroupIDRulesID structure.
type GetSecurityGroupsSecurityGroupIDRulesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityGroupsSecurityGroupIDRulesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSecurityGroupsSecurityGroupIDRulesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSecurityGroupsSecurityGroupIDRulesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSecurityGroupsSecurityGroupIDRulesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSecurityGroupsSecurityGroupIDRulesIDOK creates a GetSecurityGroupsSecurityGroupIDRulesIDOK with default headers values
func NewGetSecurityGroupsSecurityGroupIDRulesIDOK() *GetSecurityGroupsSecurityGroupIDRulesIDOK {
	return &GetSecurityGroupsSecurityGroupIDRulesIDOK{}
}

/*GetSecurityGroupsSecurityGroupIDRulesIDOK handles this case with default header values.

dummy
*/
type GetSecurityGroupsSecurityGroupIDRulesIDOK struct {
	Payload *models.SecurityGroupRule
}

func (o *GetSecurityGroupsSecurityGroupIDRulesIDOK) Error() string {
	return fmt.Sprintf("[GET /security_groups/{security_group_id}/rules/{id}][%d] getSecurityGroupsSecurityGroupIdRulesIdOK  %+v", 200, o.Payload)
}

func (o *GetSecurityGroupsSecurityGroupIDRulesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroupRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupsSecurityGroupIDRulesIDNotFound creates a GetSecurityGroupsSecurityGroupIDRulesIDNotFound with default headers values
func NewGetSecurityGroupsSecurityGroupIDRulesIDNotFound() *GetSecurityGroupsSecurityGroupIDRulesIDNotFound {
	return &GetSecurityGroupsSecurityGroupIDRulesIDNotFound{}
}

/*GetSecurityGroupsSecurityGroupIDRulesIDNotFound handles this case with default header values.

error
*/
type GetSecurityGroupsSecurityGroupIDRulesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetSecurityGroupsSecurityGroupIDRulesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /security_groups/{security_group_id}/rules/{id}][%d] getSecurityGroupsSecurityGroupIdRulesIdNotFound  %+v", 404, o.Payload)
}

func (o *GetSecurityGroupsSecurityGroupIDRulesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupsSecurityGroupIDRulesIDInternalServerError creates a GetSecurityGroupsSecurityGroupIDRulesIDInternalServerError with default headers values
func NewGetSecurityGroupsSecurityGroupIDRulesIDInternalServerError() *GetSecurityGroupsSecurityGroupIDRulesIDInternalServerError {
	return &GetSecurityGroupsSecurityGroupIDRulesIDInternalServerError{}
}

/*GetSecurityGroupsSecurityGroupIDRulesIDInternalServerError handles this case with default header values.

error
*/
type GetSecurityGroupsSecurityGroupIDRulesIDInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetSecurityGroupsSecurityGroupIDRulesIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /security_groups/{security_group_id}/rules/{id}][%d] getSecurityGroupsSecurityGroupIdRulesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSecurityGroupsSecurityGroupIDRulesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
