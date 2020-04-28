package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// Network defines the configuration structure for the `network` property
// in the configuration file
type Network struct {
	Name  string `json:"name" yaml:"name"`
	Check Check  `json:"check" yaml:"check"`
}

// Check represents a network check
type Check struct {
	Method string `json:"method" yaml:"method"`
	URL    string `json:"url" yaml:"url"`
	// StatusCode should contain the expected http status code, if not defined,
	// status codes starting with 1xx, 2xx, and 3xx will be considered successful,
	// and 4xx, and 5xx codes will be considered failures
	StatusCode int `json:"statusCode" yaml:"statusCode"`
	// Headers contains headers key-value pairs that should be present in the
	// HTTP headers of the response
	Headers map[string]string `json:"headers" yaml:"headers"`
	// ResponseBody is a regex-supported match with the response body
	ResponseBody string `json:"responseBody" yaml:"responseBody"`
	observed     *http.Response
}

func (c *Check) Run() error {
	var err error
	client := http.Client{
		Timeout: time.Second * 10,
	}
	request, err := http.NewRequest(c.Method, c.URL, nil)
	if err != nil {
		return err
	}
	c.observed, err = client.Do(request)
	if err != nil {
		return err
	}
	return nil
}

func (c Check) Verify() error {
	defer func() {
		if c.observed != nil {
			c.observed.Body.Close()
		}
	}()
	switch true {
	case c.observed == nil:
		return fmt.Errorf("Run() has not be executed")
	case c.StatusCode == 0 && c.observed.StatusCode > 399:
		return fmt.Errorf("observed status code, %v, was a non-success response (if this is expected, set StatusCode to this value)", c.observed.StatusCode)
	case c.StatusCode != 0 && c.StatusCode != c.observed.StatusCode:
		return fmt.Errorf("expected status code %v but response had status code %v", c.StatusCode, c.observed.StatusCode)
	case !c.headersMatch():
		expectedHeaders, _ := json.Marshal(c.Headers)
		responseHeaders := map[string]string{}
		for key, value := range c.observed.Header {
			responseHeaders[key] = strings.Join(value, ",")
		}
		observedHeaders, _ := json.Marshal(responseHeaders)
		return fmt.Errorf("expected headers '%s' but got '%s'", string(expectedHeaders), string(observedHeaders))
	case !c.bodiesMatch():
		observedBody, _ := ioutil.ReadAll(c.observed.Body)
		return fmt.Errorf("expected body response '%s' to match '%s' but it did not", c.ResponseBody, string(observedBody))
	}
	return nil
}

func (c Check) bodiesMatch() bool {
	body, _ := ioutil.ReadAll(c.observed.Body)
	regex := regexp.MustCompile(c.ResponseBody)
	return regex.Match(body)
}

func (c Check) headersMatch() bool {
	for key, value := range c.Headers {
		if c.observed.Header.Get(key) != value {
			return false
		}
	}
	return true
}
