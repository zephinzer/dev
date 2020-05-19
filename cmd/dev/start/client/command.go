package client

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/db"
	"github.com/zephinzer/dev/internal/gitlab"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/network"
	"github.com/zephinzer/dev/internal/notifications"
	"github.com/zephinzer/dev/internal/pivotaltracker"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.ClientCanonicalNoun,
		Aliases: constants.ClientAliases,
		Short:   "starts the dev client as a background process to provide notifications",
		Run: func(command *cobra.Command, _ []string) {
			log.Info("starting dev client...")

			log.Debug("initialising database connection...")
			connection, newConnectionError := db.NewConnection("./dev.db")
			if newConnectionError != nil {
				log.Errorf("failed to open connection to database: %s", newConnectionError)
				os.Exit(1)
			}

			networkConnectionWatcherInterval := time.Second * 20
			log.Infof("starting network connections watcher... (interval: %v)", networkConnectionWatcherInterval)
			stopNetworkConnectionWatcher := make(chan struct{})
			networkConnectionWatcher := network.WatchConnections(
				config.Global.Networks,
				networkConnectionWatcherInterval,
				stopNetworkConnectionWatcher,
			)

			pivotalTrackerNotificationWatcherInterval := time.Second * 20
			log.Infof("starting pivotal notifications watcher... (interval: %v)", pivotalTrackerNotificationWatcherInterval)
			stopPivotalWatcher := make(chan struct{})
			pivotalWatcher := pivotaltracker.WatchNotifications(
				config.Global.Platforms.PivotalTracker.AccessToken,
				config.Global.Platforms.PivotalTracker.Projects,
				connection,
				pivotalTrackerNotificationWatcherInterval,
				stopPivotalWatcher,
			)

			gitlabNotificationWatcherInterval := time.Second * 5
			log.Infof("starting gitlab notifications watcher... (interval: %v)", gitlabNotificationWatcherInterval)
			stopGitlabWatcher := make(chan struct{})
			gitlabWatcher := gitlab.WatchNotifications(
				config.Global.Platforms.Gitlab.Accounts,
				connection,
				gitlabNotificationWatcherInterval,
				stopGitlabWatcher,
			)

			close := make(chan struct{}, 1)
			go func() {
				for {
					select {
					case notification := <-pivotalWatcher:
						notifications.TriggerDesktop(notification.GetTitle(), notification.GetMessage())
					case notification := <-networkConnectionWatcher:
						notifications.TriggerDesktop(notification.GetTitle(), notification.GetMessage())
					case notification := <-gitlabWatcher:
						notifications.TriggerDesktop(notification.GetTitle(), notification.GetMessage())
					}
				}
			}()
			<-close
		},
	}
	return &cmd
}
