// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetIkePoliciesIDReader is a Reader for the GetIkePoliciesID structure.
type GetIkePoliciesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIkePoliciesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIkePoliciesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetIkePoliciesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIkePoliciesIDOK creates a GetIkePoliciesIDOK with default headers values
func NewGetIkePoliciesIDOK() *GetIkePoliciesIDOK {
	return &GetIkePoliciesIDOK{}
}

/*GetIkePoliciesIDOK handles this case with default header values.

The IKE policy was retrieved successfully.
*/
type GetIkePoliciesIDOK struct {
	Payload *models.IKEPolicy
}

func (o *GetIkePoliciesIDOK) Error() string {
	return fmt.Sprintf("[GET /ike_policies/{id}][%d] getIkePoliciesIdOK  %+v", 200, o.Payload)
}

func (o *GetIkePoliciesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IKEPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIkePoliciesIDNotFound creates a GetIkePoliciesIDNotFound with default headers values
func NewGetIkePoliciesIDNotFound() *GetIkePoliciesIDNotFound {
	return &GetIkePoliciesIDNotFound{}
}

/*GetIkePoliciesIDNotFound handles this case with default header values.

An IKE policy with the specified identifier could not be found.
*/
type GetIkePoliciesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetIkePoliciesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /ike_policies/{id}][%d] getIkePoliciesIdNotFound  %+v", 404, o.Payload)
}

func (o *GetIkePoliciesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
