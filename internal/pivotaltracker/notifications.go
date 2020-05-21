package pivotaltracker

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/types"
	pkgpivotaltracker "github.com/zephinzer/dev/pkg/pivotaltracker"
)

// GetNotifications is a wrapper around the `pkg/pivotaltracker` package that
// returns notifications as defined by types.Notifications instead of the raw
// API response
func GetNotifications(accessToken string, since ...time.Time) (types.Notifications, error) {
	ptNotifications, getNotificationsError := pkgpivotaltracker.GetNotifs(accessToken, since...)
	if getNotificationsError != nil {
		return nil, getNotificationsError
	}
	notifications := []types.Notification{}
	for _, ptNotification := range *ptNotifications {
		notifications = append(notifications, Notification(ptNotification))
	}
	return notifications, nil
}

// Notification implements types.Notification and is essentially
// a serializer of the raw API response by pivotal tracker
type Notification pkgpivotaltracker.APINotification

// GetMessage implements the types.Notification interface and returns
// the logical message component of this Notification instance
func (n Notification) GetMessage() string {
	projectName := n.Project.Name
	message := n.Message
	referenceType := n.Story.Kind
	referenceLabel := n.Story.Name
	referenceID := n.Story.ID
	datetime := n.UpdatedAt
	timestamp, err := time.Parse(constants.PivotalTrackerAPITimeFormat, datetime)
	if err == nil {
		datetime = humanize.Time(timestamp)
	}

	return fmt.Sprintf(
		"%s about %s in the %s, '%s', from %s, check it out at: https://www.pivotaltracker.com/story/show/%v",
		message, datetime, referenceType, referenceLabel, projectName, referenceID)
}

// GetTitle implements the types.Notification interface and returns
// the logical title component of this Notification instance
func (n Notification) GetTitle() string {
	datetime := n.UpdatedAt
	timestamp, err := time.Parse(constants.PivotalTrackerAPITimeFormat, datetime)
	if err == nil {
		datetime = humanize.Time(timestamp)
	}
	return fmt.Sprintf("%s about %s", n.Message, datetime)
}
