package gitlab

import (
	"database/sql"
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
							log.Tracef("['%s'@'%s'] skipping: no access token found", name, hostname)
							return
						}

						log.Tracef("['%s'@'%s'] checking for todos from gitlab...", name, hostname)
						todos, getTodosError := pkggitlab.GetTodos(hostname, accessToken, since)
						if getTodosError != nil {
							log.Warnf("['%s'@'%s'] failed to retrieve gitlab todos: %s", name, hostname, getTodosError)
							return
						}
						log.Debugf("['%s'@'%s'] received %v gitlab todos", name, hostname, len(*todos))

						for _, todo := range *todos {
							waiter.Add(1)
							go func(notif pkggitlab.APIv4Todo) {
								defer waiter.Done()
								log.Tracef("processing gitlab todo with id %v...", notif.ID)
								exists, queryExistsError := QueryNotification(todo, hostname, databaseConnection)
								if queryExistsError != nil {
									log.Warnf("failed to check existence of gitlab notification with id '%v': %s", notif.ID, queryExistsError)
									return
								}
								if !exists {
									log.Tracef("saving gitlab todo with id '%v' to the database...", todo.ID)
									if insertError := InsertNotification(notif, hostname, databaseConnection); insertError != nil {
										log.Warnf("failed to insert gitlab notification with id '%v' to the database: %s", notif.ID, insertError)
										return
									}
									log.Debugf("sending gitlab todo with id '%v' to the notifications channel", todo.ID)
									notificationsChannel <- TodoSerializer(todo)
									return
								}
								log.Debugf("skipped gitlab todo with id '%v' because it already exists in the database", todo.ID)
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
