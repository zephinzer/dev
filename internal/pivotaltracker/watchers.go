package pivotaltracker

import (
	"database/sql"
	"sync"
	"time"

	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/types"
	pkgpivotaltracker "github.com/zephinzer/dev/pkg/pivotaltracker"
)

func WatchNotifications(
	accessToken string,
	fromProjects Projects,
	databaseConnection *sql.DB,
	updateInterval time.Duration,
	stop chan struct{},
) chan types.Notification {
	notificationsChannel := make(chan types.Notification, 16)
	go func(tick <-chan time.Time) {
		for {
			select {
			case <-tick:
				log.Trace("pivotal tracker notifications watcher triggered")
				since := time.Now().UTC().Add(time.Duration(-1) * (time.Hour * 24 * 7))

				accessTokens := []string{accessToken}
				for _, project := range fromProjects {
					if len(project.AccessToken) > 0 && project.AccessToken != accessToken {
						accessTokens = append(accessTokens, project.AccessToken)
					}
				}
				log.Debugf("retrieving notifications for %v account(s)", len(accessTokens))

				var waiter sync.WaitGroup
				for _, token := range accessTokens {
					notifs, getNotifsError := pkgpivotaltracker.GetNotifs(token, since)
					if getNotifsError != nil {
						log.Warnf("failed to get notifications: %s", getNotifsError)
						continue
					}
					currentNotifications := []pkgpivotaltracker.APINotification(*notifs)
					log.Debugf("received %v notification(s) from pivotal tracker api", len(currentNotifications))
					for _, currentNotification := range currentNotifications {
						waiter.Add(1)
						go func(notif pkgpivotaltracker.APINotification) {
							defer waiter.Done()
							log.Tracef("processing pivotal tracker notification with id %v", notif.ID)
							exists, queryExistsError := QueryNotification(notif, databaseConnection)
							if queryExistsError != nil {
								log.Warnf("failed to check existence of pivotal tracker notification with id '%v': %s", notif.ID, queryExistsError)
								return
							}
							if !exists {
								log.Tracef("saving pivotal tracker notification with id %v to the database", notif.ID)
								if insertError := InsertNotification(notif, databaseConnection); insertError != nil {
									log.Warnf("failed to insert pivotal tracker notification with id '%v' to data storage: %s", notif.ID, insertError)
									return
								}
								log.Debugf("sending pivotal tracker notification with id '%v' to the notifications channel", notif.ID)
								notificationsChannel <- Notification(notif)
								return
							}
							log.Debugf("skipped pivotal tracker notification with id '%v' because it already exists in the database", notif.ID)
						}(currentNotification)
					}
				}
				waiter.Wait()
			case <-stop:
				close(notificationsChannel)
				return
			}
		}
	}(time.Tick(updateInterval))
	return notificationsChannel
}
