package oauth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type GrantRequestKey string

const (
	GrantRequestClientID     GrantRequestKey = "grant_request_client_id_key"
	GrantRequestClientSecret GrantRequestKey = "grant_request_client_secret_key"
	GrantRequestCode         GrantRequestKey = "grant_request_code_key"
	GrantRequestRedirectURI  GrantRequestKey = "grant_request_redirect_uri_key"
	GrantRequestState        GrantRequestKey = "grant_request_state_key"
)

var (
	DefaultGrantRequestMapping map[GrantRequestKey]string = map[GrantRequestKey]string{
		GrantRequestClientID:     "client_id",
		GrantRequestClientSecret: "client_secret",
		GrantRequestCode:         "code",
		GrantRequestRedirectURI:  "redirect_uri",
		GrantRequestState:        "state",
	}
)

type GrantRequest struct {
	// BaseURL defines the URL to query to grant us an access token for the authorization code we possess
	BaseURL string
	// ClientID is the ID of the client authorizing with the resource owner
	ClientID string
	// ClientSecret is the client secret for exchanging an authorization code for an access token
	ClientSecret string
	// COde is the
	Code string
	// Method is the HTTP method to use
	Method string
	// RedirectURI is the location to redirect to that can receive the authorization response
	RedirectURI string
	// State is the unique identifier for the authorization request
	State string
}

func (r GrantRequest) Do(mapping ...map[GrantRequestKey]string) (*GrantResponse, error) {
	client := http.Client{Timeout: time.Second * 5}
	request, getRequestError := r.GetRequest(mapping...)
	if getRequestError != nil {
		return nil, fmt.Errorf("failed to create request: %s", getRequestError)
	}
	response, doError := client.Do(request)
	if doError != nil {
		return nil, fmt.Errorf("failed to execute request: %s", doError)
	}
	responseBody, readAllError := ioutil.ReadAll(response.Body)
	if readAllError != nil {
		return nil, fmt.Errorf("failed to read the response body: %s", readAllError)
	}
	var grantResponse GrantResponse
	if loadFromJSONError := grantResponse.LoadFromJSON(responseBody); loadFromJSONError != nil {
		return nil, fmt.Errorf("failed to load response into an expected structure (raw response: '%s'): %s", string(responseBody), loadFromJSONError)
	}
	return &grantResponse, nil
}

func (r GrantRequest) GetRequest(mapping ...map[GrantRequestKey]string) (*http.Request, error) {
	requestFieldMap := DefaultGrantRequestMapping
	if len(mapping) > 0 {
		requestFieldMap = mapping[0]
	}
	if len(r.Method) == 0 {
		r.Method = http.MethodPost
	}
	request, newRequestError := http.NewRequest(r.Method, r.BaseURL, nil)
	if newRequestError != nil {
		return nil, fmt.Errorf("failed to instantiate http.Request for '%s %s': %s", r.Method, r.BaseURL, newRequestError)
	}
	request.Header.Add("Accept", "application/json")
	q := request.URL.Query()
	q.Add(requestFieldMap[GrantRequestClientID], r.ClientID)
	q.Add(requestFieldMap[GrantRequestClientSecret], r.ClientSecret)
	q.Add(requestFieldMap[GrantRequestCode], r.Code)
	q.Add(requestFieldMap[GrantRequestRedirectURI], r.RedirectURI)
	q.Add(requestFieldMap[GrantRequestState], r.State)
	request.URL.RawQuery = q.Encode()
	return request, nil
}
