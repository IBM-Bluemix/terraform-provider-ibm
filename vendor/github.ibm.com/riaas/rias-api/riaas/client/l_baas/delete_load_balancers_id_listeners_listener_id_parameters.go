// Code generated by go-swagger; DO NOT EDIT.

package l_baas

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

// NewDeleteLoadBalancersIDListenersListenerIDParams creates a new DeleteLoadBalancersIDListenersListenerIDParams object
// with the default values initialized.
func NewDeleteLoadBalancersIDListenersListenerIDParams() *DeleteLoadBalancersIDListenersListenerIDParams {
	var ()
	return &DeleteLoadBalancersIDListenersListenerIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLoadBalancersIDListenersListenerIDParamsWithTimeout creates a new DeleteLoadBalancersIDListenersListenerIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLoadBalancersIDListenersListenerIDParamsWithTimeout(timeout time.Duration) *DeleteLoadBalancersIDListenersListenerIDParams {
	var ()
	return &DeleteLoadBalancersIDListenersListenerIDParams{

		timeout: timeout,
	}
}

// NewDeleteLoadBalancersIDListenersListenerIDParamsWithContext creates a new DeleteLoadBalancersIDListenersListenerIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLoadBalancersIDListenersListenerIDParamsWithContext(ctx context.Context) *DeleteLoadBalancersIDListenersListenerIDParams {
	var ()
	return &DeleteLoadBalancersIDListenersListenerIDParams{

		Context: ctx,
	}
}

// NewDeleteLoadBalancersIDListenersListenerIDParamsWithHTTPClient creates a new DeleteLoadBalancersIDListenersListenerIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLoadBalancersIDListenersListenerIDParamsWithHTTPClient(client *http.Client) *DeleteLoadBalancersIDListenersListenerIDParams {
	var ()
	return &DeleteLoadBalancersIDListenersListenerIDParams{
		HTTPClient: client,
	}
}

/*DeleteLoadBalancersIDListenersListenerIDParams contains all the parameters to send to the API endpoint
for the delete load balancers ID listeners listener ID operation typically these are written to a http.Request
*/
type DeleteLoadBalancersIDListenersListenerIDParams struct {

	/*ID
	  The load balancer identifier

	*/
	ID string
	/*ListenerID
	  The listener identifier

	*/
	ListenerID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) WithTimeout(timeout time.Duration) *DeleteLoadBalancersIDListenersListenerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) WithContext(ctx context.Context) *DeleteLoadBalancersIDListenersListenerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) WithHTTPClient(client *http.Client) *DeleteLoadBalancersIDListenersListenerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) WithID(id string) *DeleteLoadBalancersIDListenersListenerIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) SetID(id string) {
	o.ID = id
}

// WithListenerID adds the listenerID to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) WithListenerID(listenerID string) *DeleteLoadBalancersIDListenersListenerIDParams {
	o.SetListenerID(listenerID)
	return o
}

// SetListenerID adds the listenerId to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) SetListenerID(listenerID string) {
	o.ListenerID = listenerID
}

// WithVersion adds the version to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) WithVersion(version string) *DeleteLoadBalancersIDListenersListenerIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete load balancers ID listeners listener ID params
func (o *DeleteLoadBalancersIDListenersListenerIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLoadBalancersIDListenersListenerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param listener_id
	if err := r.SetPathParam("listener_id", o.ListenerID); err != nil {
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
