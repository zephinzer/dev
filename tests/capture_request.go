package tests

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"

	"github.com/zephinzer/dev/pkg/utils/request"
)

type HTTPClientConsumer func(request.Doer) error
type HTTPRequestAsserter func(*http.Request) error

func HTTPRequestAsserterNoOp(_ *http.Request) error {
	return nil
}

func CaptureRequestWithTLS(caller HTTPClientConsumer, asserter HTTPRequestAsserter, response ...[]byte) error {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		asserter(r)
		if len(response) > 0 {
			w.Write(response[0])
			return
		}
		w.Write([]byte("ok"))
	}))
	client := server.Client()
	client.Transport = &http.Transport{
		DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
			return net.Dial(network, server.Listener.Addr().String())
		},
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return caller(client)
}
