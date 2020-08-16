package github

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/zephinzer/dev/internal/types"
	pkg "github.com/zephinzer/dev/pkg/github"
	"github.com/zephinzer/dev/pkg/utils/request"
)

// GetNotifications returns the serialized version of the Github notifications
func GetNotifications(client request.Doer, accessToken string, since ...time.Time) (types.Notifications, error) {
	githubNotifications, getNotificationsError := pkg.GetNotifications(client, accessToken, since...)
	if getNotificationsError != nil {
		return nil, fmt.Errorf("failed to retrieve notifications: %s", getNotificationsError)
	}
	notifications := []types.Notification{}
	for _, githubNotification := range *githubNotifications {
		notifications = append(notifications, Notification(githubNotification))
	}
	return notifications, nil
}

type Notification pkg.APIv3Notification

func (n Notification) GetMessage() string {
	title := n.GetTitle()
	itemURL := strings.Replace(n.Subject.URL, "api.github.com/repos", "github.com", 1)
	humanizedTime := humanize.Time(n.UpdatedAt)
	return fmt.Sprintf("%s (%s '%s') about %s, check it out at: %s", title, n.Subject.Type, n.Subject.Title, humanizedTime, itemURL)
}

func (n Notification) GetTitle() string {
	switch n.Reason {
	case "assign":
		return fmt.Sprintf("You were assigned to an issue")
	case "author":
		return fmt.Sprintf("Activity on a thread you created")
	case "comment":
		return fmt.Sprintf("Activity on a thread you commented on")
	case "invitation":
		return fmt.Sprintf("A repository you were invited to has activity")
	case "manual":
		return fmt.Sprintf("Activity on a thread you subscribed to")
	case "mention":
		return fmt.Sprintf("You were specifically @mentioned")
	case "review_requested":
		return fmt.Sprintf("Request to review a pull request")
	case "security_alert":
		return fmt.Sprintf("Security vulnerability discovered in your repository")
	case "state_change":
		return fmt.Sprintf("A thread state was changed")
	case "subscribed":
		return fmt.Sprintf("A repository you're watching has activity")
	case "team_mention":
		return fmt.Sprintf("Your team was @mentioned")
	}
	return fmt.Sprintf("Unknown notification type: %s", n.Reason)
}
