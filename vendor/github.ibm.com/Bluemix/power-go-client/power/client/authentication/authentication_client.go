// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new authentication API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ServiceBrokerAuthCallback returns an access token and set cookie
*/
func (a *Client) ServiceBrokerAuthCallback(params *ServiceBrokerAuthCallbackParams) (*ServiceBrokerAuthCallbackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerAuthCallbackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.auth.callback",
		Method:             "GET",
		PathPattern:        "/auth/v1/callback",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerAuthCallbackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerAuthCallbackOK), nil

}

/*
ServiceBrokerAuthDeviceCodePost requests a authorization device code
*/
func (a *Client) ServiceBrokerAuthDeviceCodePost(params *ServiceBrokerAuthDeviceCodePostParams) (*ServiceBrokerAuthDeviceCodePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerAuthDeviceCodePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.auth.device.code.post",
		Method:             "POST",
		PathPattern:        "/auth/v1/device/code",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerAuthDeviceCodePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerAuthDeviceCodePostOK), nil

}

/*
ServiceBrokerAuthDeviceTokenPost polls for authorization device token
*/
func (a *Client) ServiceBrokerAuthDeviceTokenPost(params *ServiceBrokerAuthDeviceTokenPostParams) (*ServiceBrokerAuthDeviceTokenPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerAuthDeviceTokenPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.auth.device.token.post",
		Method:             "POST",
		PathPattern:        "/auth/v1/device/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerAuthDeviceTokenPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerAuthDeviceTokenPostOK), nil

}

/*
ServiceBrokerAuthInfoToken information about current access token
*/
func (a *Client) ServiceBrokerAuthInfoToken(params *ServiceBrokerAuthInfoTokenParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBrokerAuthInfoTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerAuthInfoTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.auth.info.token",
		Method:             "GET",
		PathPattern:        "/auth/v1/info/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerAuthInfoTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerAuthInfoTokenOK), nil

}

/*
ServiceBrokerAuthInfoUser information about current user
*/
func (a *Client) ServiceBrokerAuthInfoUser(params *ServiceBrokerAuthInfoUserParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBrokerAuthInfoUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerAuthInfoUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.auth.info.user",
		Method:             "GET",
		PathPattern:        "/auth/v1/info/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerAuthInfoUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerAuthInfoUserOK), nil

}

/*
ServiceBrokerAuthLogin logins
*/
func (a *Client) ServiceBrokerAuthLogin(params *ServiceBrokerAuthLoginParams) (*ServiceBrokerAuthLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerAuthLoginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.auth.login",
		Method:             "GET",
		PathPattern:        "/auth/v1/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerAuthLoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerAuthLoginOK), nil

}

/*
ServiceBrokerAuthLogout logouts
*/
func (a *Client) ServiceBrokerAuthLogout(params *ServiceBrokerAuthLogoutParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBrokerAuthLogoutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerAuthLogoutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.auth.logout",
		Method:             "GET",
		PathPattern:        "/auth/v1/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerAuthLogoutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerAuthLogoutOK), nil

}

/*
ServiceBrokerAuthRegistration registrations of a new tenant and login
*/
func (a *Client) ServiceBrokerAuthRegistration(params *ServiceBrokerAuthRegistrationParams) (*ServiceBrokerAuthRegistrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerAuthRegistrationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.auth.registration",
		Method:             "GET",
		PathPattern:        "/auth/v1/registration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerAuthRegistrationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerAuthRegistrationOK), nil

}

/*
ServiceBrokerAuthRegistrationCallback associates the user with a tenant and returns an access token
*/
func (a *Client) ServiceBrokerAuthRegistrationCallback(params *ServiceBrokerAuthRegistrationCallbackParams) (*ServiceBrokerAuthRegistrationCallbackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerAuthRegistrationCallbackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.auth.registration.callback",
		Method:             "GET",
		PathPattern:        "/auth/v1/callback-registration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerAuthRegistrationCallbackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerAuthRegistrationCallbackOK), nil

}

/*
ServiceBrokerAuthTokenPost requests a new token from a refresh token
*/
func (a *Client) ServiceBrokerAuthTokenPost(params *ServiceBrokerAuthTokenPostParams) (*ServiceBrokerAuthTokenPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerAuthTokenPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.auth.token.post",
		Method:             "POST",
		PathPattern:        "/auth/v1/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerAuthTokenPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerAuthTokenPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
