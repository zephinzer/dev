package client

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/db"
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
			log.Debug("initialising database connection...")
			connection, newConnectionError := db.NewConnection("./dev.db")
			if newConnectionError != nil {
				log.Errorf("failed to open connection to database: %s", newConnectionError)
				os.Exit(1)
			}

			log.Debug("starting network connections watcher...")
			stopNetworkConnectionWatcher := make(chan struct{})
			networkConnectionWatcher := network.WatchConnections(
				config.Global.Networks,
				time.Second*5,
				stopNetworkConnectionWatcher,
			)

			log.Debug("starting pivotal notifications watcher...")
			stopPivotalWatcher := make(chan struct{})
			pivotalWatcher := pivotaltracker.WatchNotifications(
				config.Global.Platforms.PivotalTracker.AccessToken,
				config.Global.Platforms.PivotalTracker.Projects,
				connection,
				time.Second*15,
				stopPivotalWatcher,
			)
			go func() {
				for {
					select {
					case notification := <-pivotalWatcher:
						notifications.TriggerDesktop(notification.GetTitle(), notification.GetMessage())
					case notification := <-networkConnectionWatcher:
						notifications.TriggerDesktop(notification.GetTitle(), notification.GetMessage())
					}
				}
			}()
		},
	}
	return &cmd
}
