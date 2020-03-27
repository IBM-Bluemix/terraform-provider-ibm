// Code generated by go-swagger; DO NOT EDIT.

package hardware_platforms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewServiceBrokerHardwareplatformsGetParams creates a new ServiceBrokerHardwareplatformsGetParams object
// with the default values initialized.
func NewServiceBrokerHardwareplatformsGetParams() *ServiceBrokerHardwareplatformsGetParams {
	var ()
	return &ServiceBrokerHardwareplatformsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBrokerHardwareplatformsGetParamsWithTimeout creates a new ServiceBrokerHardwareplatformsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceBrokerHardwareplatformsGetParamsWithTimeout(timeout time.Duration) *ServiceBrokerHardwareplatformsGetParams {
	var ()
	return &ServiceBrokerHardwareplatformsGetParams{

		timeout: timeout,
	}
}

// NewServiceBrokerHardwareplatformsGetParamsWithContext creates a new ServiceBrokerHardwareplatformsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceBrokerHardwareplatformsGetParamsWithContext(ctx context.Context) *ServiceBrokerHardwareplatformsGetParams {
	var ()
	return &ServiceBrokerHardwareplatformsGetParams{

		Context: ctx,
	}
}

// NewServiceBrokerHardwareplatformsGetParamsWithHTTPClient creates a new ServiceBrokerHardwareplatformsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceBrokerHardwareplatformsGetParamsWithHTTPClient(client *http.Client) *ServiceBrokerHardwareplatformsGetParams {
	var ()
	return &ServiceBrokerHardwareplatformsGetParams{
		HTTPClient: client,
	}
}

/*ServiceBrokerHardwareplatformsGetParams contains all the parameters to send to the API endpoint
for the service broker hardwareplatforms get operation typically these are written to a http.Request
*/
type ServiceBrokerHardwareplatformsGetParams struct {

	/*RegionZone
	  The region zone of the cloud instance

	*/
	RegionZone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service broker hardwareplatforms get params
func (o *ServiceBrokerHardwareplatformsGetParams) WithTimeout(timeout time.Duration) *ServiceBrokerHardwareplatformsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service broker hardwareplatforms get params
func (o *ServiceBrokerHardwareplatformsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service broker hardwareplatforms get params
func (o *ServiceBrokerHardwareplatformsGetParams) WithContext(ctx context.Context) *ServiceBrokerHardwareplatformsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service broker hardwareplatforms get params
func (o *ServiceBrokerHardwareplatformsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service broker hardwareplatforms get params
func (o *ServiceBrokerHardwareplatformsGetParams) WithHTTPClient(client *http.Client) *ServiceBrokerHardwareplatformsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service broker hardwareplatforms get params
func (o *ServiceBrokerHardwareplatformsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegionZone adds the regionZone to the service broker hardwareplatforms get params
func (o *ServiceBrokerHardwareplatformsGetParams) WithRegionZone(regionZone *string) *ServiceBrokerHardwareplatformsGetParams {
	o.SetRegionZone(regionZone)
	return o
}

// SetRegionZone adds the regionZone to the service broker hardwareplatforms get params
func (o *ServiceBrokerHardwareplatformsGetParams) SetRegionZone(regionZone *string) {
	o.RegionZone = regionZone
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBrokerHardwareplatformsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RegionZone != nil {

		// query param regionZone
		var qrRegionZone string
		if o.RegionZone != nil {
			qrRegionZone = *o.RegionZone
		}
		qRegionZone := qrRegionZone
		if qRegionZone != "" {
			if err := r.SetQueryParam("regionZone", qRegionZone); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
