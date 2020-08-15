package github

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/pkg/utils/request"
	"github.com/zephinzer/dev/tests"
)

type NotificationsTests struct {
	suite.Suite
}

func TestNotifications(t *testing.T) {
	suite.Run(t, &NotificationsTests{})
}

func (s NotificationsTests) Test_GetNotifications() {
	timeZeroValue := time.Time{}
	systemError := tests.CaptureRequestWithTLS(
		func(client request.Doer) error {
			_, err := GetNotifications(client, "__access_token", timeZeroValue)
			return err
		},
		func(req *http.Request) error {
			s.Equal("api.github.com", req.Host)
			s.EqualValues("application/vnd.github.v3+json", req.Header["Accept"][0])
			s.EqualValues("token __access_token", req.Header["Authorization"][0])
			s.EqualValues([]string{"true"}, req.URL.Query()["participating"])
			s.EqualValues([]string{timeZeroValue.Format(constants.GithubAPITimeFormat)}, req.URL.Query()["since"])
			return nil
		},
		[]byte("[]"),
	)
	s.Nil(systemError)
}

func (s NotificationsTests) Test_GetNotifications_withError() {
	systemError := tests.CaptureRequestWithTLS(
		func(client request.Doer) error {
			_, err := GetNotifications(client, "__access_token")
			return err
		},
		tests.HTTPRequestAsserterNoOp,
		[]byte("{}"),
	)
	s.NotNil(systemError)
	s.Contains(systemError.Error(), "failed to unmarshal response")
}
