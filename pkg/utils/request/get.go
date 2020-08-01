package request

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	GetRequestTimeout = time.Second * 10
)

type GetOptions struct {
	URL     string
	Headers map[string]string
	Queries map[string]string
}

func Get(options GetOptions) (*http.Response, error) {
	targetURL, err := url.Parse(options.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s': %s", options.URL, err)
	}
	query := targetURL.Query()
	for key, value := range options.Queries {
		query.Add(key, value)
	}
	targetURL.RawQuery = query.Encode()
	client := http.Client{Timeout: GetRequestTimeout}
	request, err := http.NewRequest("GET", targetURL.String(), nil)
	for key, value := range options.Headers {
		request.Header.Add(key, value)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}
	return client.Do(request)
}
