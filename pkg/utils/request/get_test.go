package request

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GetTests struct {
	suite.Suite
}

func TestGet(t *testing.T) {
	suite.Run(t, &GetTests{})
}

func (s *GetTests) Test_Get() {
	var observedHost, observedMethod, observedProto string
	observedHeaders := map[string][]string{}
	observedQueries := map[string][]string{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		observedHost = r.Host
		observedMethod = r.Method
		observedProto = r.Proto
		for key, value := range r.Header {
			observedHeaders[key] = value
		}
		for key, value := range r.URL.Query() {
			observedQueries[key] = value
		}
		for key, value := range r.URL.Query() {
			observedQueries[key] = value
		}
		w.Write([]byte("ok"))
	}))
	address := server.URL
	fmt.Println("the address is here")
	fmt.Println(address)
	_, err := Get(GetOptions{
		URL:     address,
		Headers: map[string]string{"hello": "world"},
		Queries: map[string]string{"hola": "mundo"},
	})
	s.Nil(err)
	if err != nil {
		return
	}
	parsedAddress, err := url.Parse(address)
	s.Nil(err)
	if err != nil {
		return
	}
	s.Equal("http/1.1", strings.ToLower(observedProto))
	s.Equal("get", strings.ToLower(observedMethod))
	s.Equal(parsedAddress.Host, observedHost)
	s.Equal([]string{"world"}, observedHeaders["Hello"])
	s.Equal(map[string][]string{"hola": {"mundo"}}, observedQueries)
}
