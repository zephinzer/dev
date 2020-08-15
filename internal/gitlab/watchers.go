package gitlab

import (
	"database/sql"
	"net/http"
	"sync"
	"time"

	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/types"
	pkggitlab "github.com/zephinzer/dev/pkg/gitlab"
)

func WatchNotifications(
	accounts []AccountConfig,
	databaseConnection *sql.DB,
	updateInterval time.Duration,
	stop chan struct{},
) chan types.Notification {
	notificationsChannel := make(chan types.Notification, 16)
	go func(tick <-chan time.Time) {
		for {
			select {
			case <-stop:
				close(notificationsChannel)
				return
			case <-tick:
				log.Trace("gitlab notifications watcher triggered")
				since := time.Now().UTC().Add(time.Duration(-1) * (time.Hour * 24 * 7))
				log.Debugf("retrieving notifications for %v accounts", len(accounts))
				var waiter sync.WaitGroup
				for _, account := range accounts {
					waiter.Add(1)
					go func(acc AccountConfig) {
						defer waiter.Done()
						name := acc.Name
						if len(name) == 0 {
							name = "unnamed-gitlab"
						}
						hostname := acc.Hostname
						if len(hostname) == 0 {
							hostname = constants.DefaultGitlabHostname
						}
						accessToken := acc.AccessToken
						if len(accessToken) == 0 {
							log.Tracef("[%s] skipping: no access token found for gitlab at %s", name, hostname)
							return
						}

						log.Tracef("[%s] checking for todos from gitlab at %s...", name, hostname)
						client := &http.Client{Timeout: constants.DefaultAPICallTimeout}
						todos, getTodosError := pkggitlab.GetTodos(client, hostname, accessToken, since)
						if getTodosError != nil {
							log.Warnf("[%s] failed to retrieve gitlab todos from %s: %s", name, hostname, getTodosError)
							return
						}
						log.Debugf("[%s] received %v todos from gitlab at %s", name, len(*todos), hostname)

						for _, todo := range *todos {
							waiter.Add(1)
							go func(notif pkggitlab.APIv4Todo) {
								defer waiter.Done()
								log.Tracef("[%s] processing gitlab todo with id '%v@%s'...", name, notif.ID, hostname)
								exists, queryExistsError := QueryNotification(todo, hostname, databaseConnection)
								if queryExistsError != nil {
									log.Warnf("[%s] failed to check existence of gitlab notification with id '%v@%s': %s", name, notif.ID, hostname, queryExistsError)
									return
								}
								if !exists {
									log.Tracef("[%s] saving gitlab todo with id '%v@%s' to the database...", name, notif.ID, hostname)
									if insertError := InsertNotification(notif, hostname, databaseConnection); insertError != nil {
										log.Warnf("failed to insert gitlab notification with id '%v@%s' to the database: %s", notif.ID, hostname, insertError)
										return
									}
									log.Debugf("[%s] triggering notification for gitlab todo with id '%v@%s'", name, notif.ID, hostname)
									notificationsChannel <- TodoSerializer(notif)
									return
								}
								log.Debugf("[%s] skipped gitlab todo with id '%v@%s' because it already exists in the database", name, notif.ID, hostname)
							}(todo)
						}
					}(account)
				}
				waiter.Wait()
			}
		}
	}(time.Tick(updateInterval))
	return notificationsChannel
}
