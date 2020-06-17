package oauth

import (
	"errors"
	"net/url"
)

type AuthorizationResponseKey string

const (
	AuthorizationResponseCode             AuthorizationResponseKey = "authorization_response_code_key"
	AuthorizationResponseError            AuthorizationResponseKey = "authorization_response_error_key"
	AuthorizationResponseErrorDescription AuthorizationResponseKey = "authorization_response_error_description_key"
	AuthorizationResponseErrorURI         AuthorizationResponseKey = "authorization_response_error_uri_key"
	AuthorizationResponseState            AuthorizationResponseKey = "authorization_response_state_key"
)

var (
	DefaultAuthorizationResponseMapping map[AuthorizationResponseKey]string = map[AuthorizationResponseKey]string{
		AuthorizationResponseCode:             "code",
		AuthorizationResponseError:            "error",
		AuthorizationResponseErrorDescription: "error_description",
		AuthorizationResponseErrorURI:         "error_uri",
		AuthorizationResponseState:            "state",
	}
)

// AuthorizationResponse represents the response sent
// by the authorization server after a user has authorized
// the application
type AuthorizationResponse struct {
	// Code is the authorization code returned by the authorization
	// server if the authorization succeeded
	Code string
	// Error is an error symbol representing the type of error
	Error error
	// ErrorDescription is an arbitrary string describing the error
	ErrorDescription string
	// ErrorURI is a URL that can assist a consumer in understanding
	// the returned ErrorDescription
	ErrorURI string
	// State is the unique string sent to the authorization server
	// for identification of this authorization request
	State string
}

// LoadFromQuery loads this instance of AuthorizationResponse using the URL query
// values from the callback response
func (r *AuthorizationResponse) LoadFromQuery(responseQuery url.Values, mapping ...map[AuthorizationResponseKey]string) error {
	responseFieldMap := DefaultAuthorizationResponseMapping
	if len(mapping) > 0 {
		responseFieldMap = mapping[0]
	}
	codeField := responseFieldMap[AuthorizationResponseCode]
	errorField := responseFieldMap[AuthorizationResponseError]
	errorDescriptionField := responseFieldMap[AuthorizationResponseErrorDescription]
	errorURIField := responseFieldMap[AuthorizationResponseErrorURI]
	stateField := responseFieldMap[AuthorizationResponseState]

	err := responseQuery.Get(errorField)
	if len(err) > 0 {
		r.Error = errors.New(err)
		r.ErrorDescription = responseQuery.Get(errorDescriptionField)
		r.ErrorURI = responseQuery.Get(errorURIField)
		return r.Error
	}
	r.Code = responseQuery.Get(codeField)
	r.State = responseQuery.Get(stateField)
	return nil
}
