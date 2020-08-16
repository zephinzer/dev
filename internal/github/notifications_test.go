package github

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	pkg "github.com/zephinzer/dev/pkg/github"
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
	s.Nil(tests.CaptureRequestWithTLS(
		func(client request.Doer) error {
			_, err := GetNotifications(client, "__access_token")
			return err
		},
		func(req *http.Request) error {
			s.Equal([]string{"token __access_token"}, req.Header["Authorization"])
			return nil
		},
		[]byte("[]"),
	))
}

func (s NotificationsTests) Test_Notification_GetMessage() {
	notif := Notification{
		Subject: pkg.APIv3NotificationSubject{
			Title: "__title",
			Type:  "__type",
			URL:   "api.github.com/repos/__url",
		},
	}
	message := notif.GetMessage()
	s.Contains(message, "__title", "should contain the title")
	s.Contains(message, "__type", "should contain the type")
	s.Contains(message, "github.com/__url", "should contain a modified url")
}

func (s NotificationsTests) Test_Notification_GetTitle() {
	notifs := []Notification{
		{Reason: "author"},
		{Reason: "comment"},
		{Reason: "invitation"},
		{Reason: "manual"},
		{Reason: "mention"},
		{Reason: "review_requested"},
		{Reason: "security_alert"},
		{Reason: "state_change"},
		{Reason: "subscribed"},
		{Reason: "team_mention"},
	}
	seen := map[string]string{}
	for _, notif := range notifs {
		title := notif.GetTitle()
		val, ok := seen[title]
		s.Falsef(ok, "returned strings should be unique according to .Reason, reasons '%s' and '%s' have the same description", notif.Reason, val)
		seen[title] = notif.Reason
	}
}
