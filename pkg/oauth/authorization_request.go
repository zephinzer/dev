package oauth

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

type AuthorizationRequestKey string

const (
	AuthorizationRequestClientID    AuthorizationRequestKey = "authorization_request_client_id_key"
	AuthorizationRequestRedirectURI AuthorizationRequestKey = "authorization_request_redirect_uri_key"
	AuthorizationRequestScope       AuthorizationRequestKey = "authorization_request_scope_key"
	AuthorizationRequestState       AuthorizationRequestKey = "authorization_request_state_key"
	DefaultScopeDelimiter                                   = " "
)

var (
	DefaultAuthorizationRequestMapping map[AuthorizationRequestKey]string = map[AuthorizationRequestKey]string{
		AuthorizationRequestClientID:    "client_id",
		AuthorizationRequestRedirectURI: "redirect_uri",
		AuthorizationRequestScope:       "scope",
		AuthorizationRequestState:       "state",
	}
)

type AuthorizationRequest struct {
	BaseURL string
	// ClientID is the ID of the client authorizing with the resource owner
	ClientID string
	// RedirectURI is the location to redirect to that can receive the authorization response
	RedirectURI string
	// Scopes is a list of resource scopes to request for authorization for
	Scopes []string
	// ScopeDelimiter is used to concatenate the provided Scopes into a single string
	ScopeDelimiter string
	// State is a runtime-generated string to identify this authorization request
	State string
}

// OpenInBrowser is a convenience method to open the authorization page in the user's
// default browser
func (r *AuthorizationRequest) OpenInBrowser(mapping ...map[AuthorizationRequestKey]string) error {
	authorizationRequestURL, getURLError := r.GetURL(mapping...)
	if getURLError != nil {
		return fmt.Errorf("failed to open the authorization request url")
	}
	Open(authorizationRequestURL)
	return nil
}

// GetURL builds the resource authorization URL for the user
func (r *AuthorizationRequest) GetURL(mapping ...map[AuthorizationRequestKey]string) (string, error) {
	queryFieldMap := DefaultAuthorizationRequestMapping
	if len(mapping) > 0 {
		queryFieldMap = mapping[0]
	}
	authorizationURL, parseError := url.Parse(r.BaseURL)
	if parseError != nil {
		return "", fmt.Errorf("failed to parse the base url '%s': %s", r.BaseURL, parseError)
	}
	query := authorizationURL.Query()
	query.Add(queryFieldMap[AuthorizationRequestClientID], r.ClientID)
	query.Add(queryFieldMap[AuthorizationRequestRedirectURI], r.RedirectURI)
	scopeDelimiter := r.ScopeDelimiter
	if len(scopeDelimiter) == 0 {
		scopeDelimiter = DefaultScopeDelimiter
	}
	query.Add(queryFieldMap[AuthorizationRequestScope], strings.Join(r.Scopes, scopeDelimiter))
	if len(r.State) == 0 {
		r.State = uuid.New().String()
	}
	query.Add(queryFieldMap[AuthorizationRequestState], r.State)
	authorizationURL.RawQuery = query.Encode()
	return authorizationURL.String(), nil
}
