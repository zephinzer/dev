package pivotaltracker

import (
	"time"

	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/internal/types"
	pkgpivotaltracker "github.com/usvc/dev/pkg/pivotaltracker"
)

func WatchNotifications(
	accessToken string,
	fromProjects pkgpivotaltracker.Projects,
	updateInterval time.Duration,
	stop chan struct{},
) chan types.Notification {
	notificationsChannel := make(chan types.Notification, 1)
	go func(tick <-chan time.Time) {
		for {
			select {
			case <-tick:
				log.Trace("pivotal tracker notifications watcher triggered")
				// TODO - implement code to retrieve notifications
				notificationsChannel <- PlaceholderNotification{}
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
