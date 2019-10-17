// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSecurityGroupsIDParams creates a new DeleteSecurityGroupsIDParams object
// with the default values initialized.
func NewDeleteSecurityGroupsIDParams() *DeleteSecurityGroupsIDParams {
	var ()
	return &DeleteSecurityGroupsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSecurityGroupsIDParamsWithTimeout creates a new DeleteSecurityGroupsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSecurityGroupsIDParamsWithTimeout(timeout time.Duration) *DeleteSecurityGroupsIDParams {
	var ()
	return &DeleteSecurityGroupsIDParams{

		timeout: timeout,
	}
}

// NewDeleteSecurityGroupsIDParamsWithContext creates a new DeleteSecurityGroupsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSecurityGroupsIDParamsWithContext(ctx context.Context) *DeleteSecurityGroupsIDParams {
	var ()
	return &DeleteSecurityGroupsIDParams{

		Context: ctx,
	}
}

// NewDeleteSecurityGroupsIDParamsWithHTTPClient creates a new DeleteSecurityGroupsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSecurityGroupsIDParamsWithHTTPClient(client *http.Client) *DeleteSecurityGroupsIDParams {
	var ()
	return &DeleteSecurityGroupsIDParams{
		HTTPClient: client,
	}
}

/*DeleteSecurityGroupsIDParams contains all the parameters to send to the API endpoint
for the delete security groups ID operation typically these are written to a http.Request
*/
type DeleteSecurityGroupsIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The security group identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) WithTimeout(timeout time.Duration) *DeleteSecurityGroupsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) WithContext(ctx context.Context) *DeleteSecurityGroupsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) WithHTTPClient(client *http.Client) *DeleteSecurityGroupsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) WithGeneration(generation int64) *DeleteSecurityGroupsIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) WithID(id string) *DeleteSecurityGroupsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) WithVersion(version string) *DeleteSecurityGroupsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete security groups ID params
func (o *DeleteSecurityGroupsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSecurityGroupsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
