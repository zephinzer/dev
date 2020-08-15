package github

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AccountTests struct {
	suite.Suite
}

func TestAccount(t *testing.T) {
	suite.Run(t, &AccountTests{})
}

func (s AccountTests) Test_GetAccount() {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Equal("api.github.com", r.Host)
		s.EqualValues("application/vnd.github.v3+json", r.Header["Accept"][0])
		s.EqualValues("token __access_token", r.Header["Authorization"][0])
		w.Write([]byte("{}"))
	}))
	client := server.Client()
	client.Transport = &http.Transport{
		DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
			return net.Dial(network, server.Listener.Addr().String())
		},
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var err error
	_, err = GetAccount(client, "__access_token")
	s.Nil(err)
	if err != nil {
		return
	}
}
