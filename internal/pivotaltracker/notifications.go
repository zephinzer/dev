package pivotaltracker

import (
	"time"

	"github.com/usvc/dev/internal/types"
)

type Notifications []types.Notification

func WatchNotifications(updateInterval time.Duration, stop chan struct{}) chan Notifications {
	notificationsChannel := make(chan Notifications, 1)
	go func(tick <-chan time.Time) {
		for {
			select {
			case <-tick:
				// TODO - implement code to retrieve notifications
				notificationsChannel <- []types.Notification{PlaceholderNotification{}}
			case <-stop:
				return
			}
		}
	}(time.Tick(updateInterval))
	return notificationsChannel
}

type PlaceholderNotification struct{}

func (phnotif PlaceholderNotification) GetMessage() string {
	return "placeholder message"
}

func (phnotif PlaceholderNotification) GetTitle() string {
	return "placeholder title"
}
