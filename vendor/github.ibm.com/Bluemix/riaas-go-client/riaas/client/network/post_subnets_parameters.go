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

// NewPostSubnetsParams creates a new PostSubnetsParams object
// with the default values initialized.
func NewPostSubnetsParams() *PostSubnetsParams {
	var ()
	return &PostSubnetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSubnetsParamsWithTimeout creates a new PostSubnetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSubnetsParamsWithTimeout(timeout time.Duration) *PostSubnetsParams {
	var ()
	return &PostSubnetsParams{

		timeout: timeout,
	}
}

// NewPostSubnetsParamsWithContext creates a new PostSubnetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSubnetsParamsWithContext(ctx context.Context) *PostSubnetsParams {
	var ()
	return &PostSubnetsParams{

		Context: ctx,
	}
}

// NewPostSubnetsParamsWithHTTPClient creates a new PostSubnetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSubnetsParamsWithHTTPClient(client *http.Client) *PostSubnetsParams {
	var ()
	return &PostSubnetsParams{
		HTTPClient: client,
	}
}

/*PostSubnetsParams contains all the parameters to send to the API endpoint
for the post subnets operation typically these are written to a http.Request
*/
type PostSubnetsParams struct {

	/*Body*/
	Body PostSubnetsBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post subnets params
func (o *PostSubnetsParams) WithTimeout(timeout time.Duration) *PostSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post subnets params
func (o *PostSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post subnets params
func (o *PostSubnetsParams) WithContext(ctx context.Context) *PostSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post subnets params
func (o *PostSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post subnets params
func (o *PostSubnetsParams) WithHTTPClient(client *http.Client) *PostSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post subnets params
func (o *PostSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post subnets params
func (o *PostSubnetsParams) WithBody(body PostSubnetsBody) *PostSubnetsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post subnets params
func (o *PostSubnetsParams) SetBody(body PostSubnetsBody) {
	o.Body = body
}

// WithGeneration adds the generation to the post subnets params
func (o *PostSubnetsParams) WithGeneration(generation int64) *PostSubnetsParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the post subnets params
func (o *PostSubnetsParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithVersion adds the version to the post subnets params
func (o *PostSubnetsParams) WithVersion(version string) *PostSubnetsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post subnets params
func (o *PostSubnetsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
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
