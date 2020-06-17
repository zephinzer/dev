package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
)

type GrantResponseKey string

const (
	GrantResponseAccessToken GrantResponseKey = "grant_response_access_token_key"
	GrantResponseError       GrantResponseKey = "grant_response_error"
	GrantResponseScope       GrantResponseKey = "grant_response_scope_key"
	GrantResponseTokenType   GrantResponseKey = "grant_response_type_key"
)

var (
	DefaultGrantResponseMapping map[GrantResponseKey]string = map[GrantResponseKey]string{
		GrantResponseAccessToken: "access_token",
		GrantResponseError:       "error",
		GrantResponseScope:       "scope",
		GrantResponseTokenType:   "token_type",
	}
)

// GrantResponse represents the response to the request
// by the application to exchange an authorization code for an
// access token
type GrantResponse struct {
	AccessToken string
	Scope       string
	TokenType   string
	Error       error
}

// LoadFromJSON parses the provided JSON-formatted responseBody into a GrantResponse instance
func (r *GrantResponse) LoadFromJSON(responseBody []byte, mapping ...map[GrantResponseKey]string) error {
	responseFieldMap := DefaultGrantResponseMapping
	if len(mapping) > 0 {
		responseFieldMap = mapping[0]
	}
	var unmarshalledResponse map[string]interface{}
	unmarshalError := json.Unmarshal(responseBody, &unmarshalledResponse)
	if unmarshalError != nil {
		return fmt.Errorf("failed to parse response '%s': %s", string(responseBody), unmarshalError)
	}
	var ok bool
	var err string
	if err, ok = unmarshalledResponse[responseFieldMap[GrantResponseError]].(string); ok {
		r.Error = errors.New(err)
		return nil
	}
	if r.AccessToken, ok = unmarshalledResponse[responseFieldMap[GrantResponseAccessToken]].(string); !ok {
		return fmt.Errorf("failed to retrieve the access-token field using key '%s'", responseFieldMap[GrantResponseAccessToken])
	}
	if r.Scope, ok = unmarshalledResponse[responseFieldMap[GrantResponseScope]].(string); !ok {
		return fmt.Errorf("failed to retrieve the scope field using key '%s'", responseFieldMap[GrantResponseScope])
	}
	if r.TokenType, ok = unmarshalledResponse[responseFieldMap[GrantResponseTokenType]].(string); !ok {
		return fmt.Errorf("failed to retrieve the token-type field using key '%s'", responseFieldMap[GrantResponseTokenType])
	}
	return nil
}
