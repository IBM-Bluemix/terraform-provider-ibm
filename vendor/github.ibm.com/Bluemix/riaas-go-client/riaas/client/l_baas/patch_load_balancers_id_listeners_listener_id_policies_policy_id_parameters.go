// Code generated by go-swagger; DO NOT EDIT.

package l_baas

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

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams creates a new PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams object
// with the default values initialized.
func NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams() *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	var ()
	return &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithTimeout creates a new PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithTimeout(timeout time.Duration) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	var ()
	return &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams{

		timeout: timeout,
	}
}

// NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithContext creates a new PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithContext(ctx context.Context) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	var ()
	return &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams{

		Context: ctx,
	}
}

// NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithHTTPClient creates a new PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithHTTPClient(client *http.Client) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	var ()
	return &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams{
		HTTPClient: client,
	}
}

/*PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams contains all the parameters to send to the API endpoint
for the patch load balancers ID listeners listener ID policies policy ID operation typically these are written to a http.Request
*/
type PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams struct {

	/*Body
	  The policy template

	*/
	Body *models.ListenerPolicyTemplatePatch
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The load balancer identifier

	*/
	ID string
	/*ListenerID
	  The listener identifier

	*/
	ListenerID string
	/*PolicyID
	  The policy identifier

	*/
	PolicyID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithTimeout(timeout time.Duration) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithContext(ctx context.Context) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithHTTPClient(client *http.Client) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithBody(body *models.ListenerPolicyTemplatePatch) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetBody(body *models.ListenerPolicyTemplatePatch) {
	o.Body = body
}

// WithGeneration adds the generation to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithGeneration(generation int64) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithID(id string) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetID(id string) {
	o.ID = id
}

// WithListenerID adds the listenerID to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithListenerID(listenerID string) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetListenerID(listenerID)
	return o
}

// SetListenerID adds the listenerId to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetListenerID(listenerID string) {
	o.ListenerID = listenerID
}

// WithPolicyID adds the policyID to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithPolicyID(policyID string) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetPolicyID(policyID string) {
	o.PolicyID = policyID
}

// WithVersion adds the version to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithVersion(version string) *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch load balancers ID listeners listener ID policies policy ID params
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

	// path param listener_id
	if err := r.SetPathParam("listener_id", o.ListenerID); err != nil {
		return err
	}

	// path param policy_id
	if err := r.SetPathParam("policy_id", o.PolicyID); err != nil {
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
