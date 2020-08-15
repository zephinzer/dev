package request

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var ValidHTTPMethods = map[string]bool{
	http.MethodConnect: true,
	http.MethodDelete:  true,
	http.MethodGet:     true,
	http.MethodHead:    true,
	http.MethodOptions: true,
	http.MethodPatch:   true,
	http.MethodPost:    true,
	http.MethodPut:     true,
	http.MethodTrace:   true,
}

// CreateOptions provides the input to the Create method
type CreateOptions struct {
	Body    []byte
	Headers map[string][]string
	Method  string
	Queries map[string][]string
	URL     string
}

// Create is a convenience function to increase readability in consuming
// code that needs to create an http.Request instance for calling an external
// service
func Create(options CreateOptions) (*http.Request, error) {
	// configure the base query
	targetURL, err := url.Parse(options.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s': %s", options.URL, err)
	}

	// configure url queries
	query := targetURL.Query()
	for key, values := range options.Queries {
		for _, value := range values {
			query.Add(key, value)
		}
	}
	targetURL.RawQuery = query.Encode()

	// configure the request object
	method := strings.ToUpper(options.Method)
	if accurate, valid := ValidHTTPMethods[method]; !accurate || !valid {
		method = http.MethodGet
	}
	request, err := http.NewRequest(method, targetURL.String(), bytes.NewBuffer(options.Body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	// configure http headers
	for key, values := range options.Headers {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}

	return request, nil
}
