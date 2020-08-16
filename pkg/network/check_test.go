package network

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NetworkTests struct {
	suite.Suite
}

func TestNetwork(t *testing.T) {
	suite.Run(t, &NetworkTests{})
}

func (s *NetworkTests) Test_Check_GetObserved() {
	expectedStatus := "777 test"
	expectedStatusCode := 777
	check := Check{
		observed: &http.Response{
			Status:     expectedStatus,
			StatusCode: expectedStatusCode,
		},
	}
	observed, getObservedError := check.GetObserved()
	s.Nil(getObservedError)
	s.Equal(expectedStatus, observed.Status)
	s.Equal(expectedStatusCode, observed.StatusCode)
}

func (s *NetworkTests) Test_Check_GetObserved_checkNotRunYet() {
	check := Check{}
	observed, getObservedError := check.GetObserved()
	s.Zero(observed)
	s.NotNil(getObservedError)
	s.Contains(getObservedError.Error(), "not been run yet")
}

func (s *NetworkTests) Test_Check_Run() {
	expectedBody := "hello"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expectedBody))
	}))
	defer server.Close()
	check := Check{
		URL: server.URL,
	}
	runError := check.Run()
	s.Nil(runError)
	observed, getObservedError := check.GetObserved()
	s.Nil(getObservedError)
	if getObservedError != nil {
		return
	}
	s.NotZero(observed)
	body, readAllError := ioutil.ReadAll(observed.Body)
	s.Nil(readAllError)
	s.Equal(expectedBody, string(body))
}

func (s *NetworkTests) Test_Check_Verify_notRunYet() {
	check := Check{}
	verifyError := check.Verify()
	s.NotNil(verifyError)
}

func (s *NetworkTests) Test_Check_Verify_deferDoesNotPanic() {
	defer func() {
		s.Nil(recover())
	}()
	check := Check{
		observed: &http.Response{
			Body: nil,
		},
	}
	check.Verify()
}

func (s *NetworkTests) Test_Check_Verify_statusCodeNotSet() {
	check := Check{
		observed: &http.Response{
			StatusCode: 400,
		},
	}
	verify := check.Verify()
	s.NotNil(verify)
	s.Contains(verify.Error(), "non-success response")
	check.StatusCode = check.observed.StatusCode
	verify = check.Verify()
	s.Nil(verify)
}

func (s *NetworkTests) Test_Check_Verify_statusCodeSet() {
	check := Check{
		StatusCode: 400,
		observed: &http.Response{
			StatusCode: 400,
		},
	}
	verify := check.Verify()
	s.Nil(verify)
	check.StatusCode--
	verify = check.Verify()
	s.NotNil(verify)
}

func (s *NetworkTests) Test_Check_Verify_body() {
	expectedBody := "hello"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expectedBody))
	}))
	defer server.Close()
	check := Check{
		ResponseBody: expectedBody,
		URL:          server.URL,
	}
	runError := check.Run()
	s.Nil(runError)
	if runError != nil {
		return
	}
	verify := check.Verify()
	s.Nil(verify)
}

func (s *NetworkTests) Test_Check_Verify_bodyWrong() {
	expectedBody := "hello"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expectedBody))
	}))
	defer server.Close()
	check := Check{
		ResponseBody: expectedBody + " world",
		URL:          server.URL,
	}
	runError := check.Run()
	s.Nil(runError)
	if runError != nil {
		return
	}
	verify := check.Verify()
	s.NotNil(verify)
	s.Contains(verify.Error(), "expected body")
}

func (s *NetworkTests) Test_Check_Verify_headers() {
	check := Check{
		Headers: map[string]string{
			"hello": "world",
		},
		observed: &http.Response{
			Header: http.Header{
				http.CanonicalHeaderKey("hello"): []string{"world"},
			},
		},
	}
	verify := check.Verify()
	s.Nil(verify)
}

func (s *NetworkTests) Test_Check_Verify_headersAreWrong() {
	check := Check{
		Headers: map[string]string{
			"hello": "planet",
		},
		observed: &http.Response{
			Header: http.Header{
				http.CanonicalHeaderKey("hello"): []string{"world"},
			},
		},
	}
	verify := check.Verify()
	s.NotNil(verify)
	s.Contains(verify.Error(), "expected headers")
}
