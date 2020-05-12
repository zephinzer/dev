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

// GetObserved returns the response object from the check; if no check has
// been performed yet, the http.Response will be of zero value with an error
func (c *Check) GetObserved() (http.Response, error) {
	if c.observed == nil {
		return http.Response{}, fmt.Errorf("network check to '%s' has not been run yet", c.URL)
	}
	return *c.observed, nil
}

// Run executes the target network check and places the response in the
// .observed property (used by .Verify to check if the check succeeded); Run
// will return an error on any network issues that result in the .observed
// property not being populated
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

// Verify compares the expected values and the observed outcome, returning
// an error if the verification failed.
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

// bodiesMatch is a utilities function to check if the expected response body
// matches the observed response body
func (c Check) bodiesMatch() bool {
	body, _ := ioutil.ReadAll(c.observed.Body)
	regex := regexp.MustCompile(c.ResponseBody)
	return regex.Match(body)
}

// headersMatch is a utilities function to check if the expected headers
// match the observed headers
func (c Check) headersMatch() bool {
	for key, value := range c.Headers {
		if c.observed.Header.Get(key) != value {
			return false
		}
	}
	return true
}
