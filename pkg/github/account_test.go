package github

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/pkg/utils/request"
	"github.com/zephinzer/dev/tests"
)

type AccountTests struct {
	suite.Suite
}

func TestAccount(t *testing.T) {
	suite.Run(t, &AccountTests{})
}

func (s AccountTests) Test_GetAccount() {
	systemError := tests.CaptureRequestWithTLS(
		func(client request.Doer) error {
			_, err := GetAccount(client, "__access_token")
			return err
		},
		func(req *http.Request) error {
			s.Equal("api.github.com", req.Host)
			s.EqualValues("application/vnd.github.v3+json", req.Header["Accept"][0])
			s.EqualValues("token __access_token", req.Header["Authorization"][0])
			return nil
		},
		[]byte("{}"),
	)
	s.Nil(systemError)
}

func (s AccountTests) Test_GetAccount_withError() {
	systemError := tests.CaptureRequestWithTLS(
		func(client request.Doer) error {
			_, err := GetAccount(client, "__access_token")
			return err
		},
		tests.HTTPRequestAsserterNoOp,
		[]byte("hi"),
	)
	s.NotNil(systemError)
	s.Contains(systemError.Error(), "failed to unmarshal response")
}
