package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CreateTests struct {
	suite.Suite
}

func TestCreate(t *testing.T) {
	suite.Run(t, &CreateTests{})
}

func (s CreateTests) Test_Create_setsBody() {
	body := []byte("hello world")
	req, err := Create(CreateOptions{
		Body: body,
		URL:  "https://testurl.com/path/to/endpoint",
	})
	s.Nil(err)
	if err != nil {
		return
	}
	s.NotEqual(http.NoBody, req.Body)
}

func (s CreateTests) Test_Create_setsURL() {
	req, err := Create(CreateOptions{
		URL: "https://testurl.com/path/to/endpoint?with=queries",
	})
	s.Nil(err)
	if err != nil {
		return
	}
	q := req.URL.Query()
	s.Equal("/path/to/endpoint", req.URL.Path)
	s.Equal("testurl.com", req.Host)
	s.Equal("https", req.URL.Scheme)
	s.Equal("queries", q.Get("with"))
	s.EqualValues(http.NoBody, req.Body)
}

func (s CreateTests) Test_Create_setsHeader() {
	req, err := Create(CreateOptions{
		URL: "https://testurl.com/path/to/endpoint",
		Headers: map[string][]string{
			"header_1": {"one"},
			"header_2": {"two", "2"},
		},
	})
	s.Nil(err)
	if err != nil {
		return
	}
	observed := map[string][]string{}
	for key, values := range req.Header {
		observed[key] = values
	}
	s.EqualValues([]string{"one"}, observed["Header_1"])
	s.EqualValues([]string{"two", "2"}, observed["Header_2"])
}

func (s CreateTests) Test_Create_setsMethod() {
	req, err := Create(CreateOptions{
		Method: http.MethodConnect,
		URL:    "https://testurl.com/path/to/endpoint",
	})
	s.Nil(err)
	if err != nil {
		return
	}
	s.Equal(http.MethodConnect, req.Method)
}

func (s CreateTests) Test_Create_setsMethod_defaultsToGET() {
	req, err := Create(CreateOptions{
		Method: "asdasd",
		URL:    "https://testurl.com/path/to/endpoint",
	})
	s.Nil(err)
	if err != nil {
		return
	}
	s.Equal(http.MethodGet, req.Method)
}

func (s CreateTests) Test_Create_setsQueries() {
	req, err := Create(CreateOptions{
		URL: "https://testurl.com/path/to/endpoint",
		Queries: map[string][]string{
			"query_1": {"one"},
			"query_2": {"two", "2"},
		},
	})
	s.Nil(err)
	if err != nil {
		return
	}
	observed := map[string][]string{}
	for key, values := range req.URL.Query() {
		observed[key] = values
	}
	s.EqualValues([]string{"one"}, observed["query_1"])
	s.EqualValues([]string{"two", "2"}, observed["query_2"])
}
