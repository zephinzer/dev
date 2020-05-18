package utils

import (
	"net/http"
	"net/url"
	"time"
)

func HTTPGet(targetURL url.URL, headers map[string]string) (*http.Response, error) {
	timeout := time.Second * 10
	client := http.Client{Timeout: timeout}
	request, err := http.NewRequest("GET", targetURL.String(), nil)
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	if err != nil {
		return nil, err
	}
	return client.Do(request)
}
