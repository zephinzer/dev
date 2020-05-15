package client

import (
	"os"
	"runtime"
	"time"

	"github.com/getlantern/systray"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/db"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/internal/network"
	"github.com/usvc/dev/internal/notifications"
	"github.com/usvc/dev/internal/pivotaltracker"
	"github.com/usvc/dev/internal/systemtray"
	"github.com/usvc/dev/pkg/utils"
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

			log.Debugf("adding system tray icon...")
			systemTrayStopped := make(chan struct{})
			systemtray.Start(systemtray.Menu{
				{
					Label:   "About",
					Tooltip: "The dev repository",
					Handler: func() {
						theURL := constants.CanonicalRepositoryURL
						log.Infof("opening '%s' for the '%s' platform", theURL, runtime.GOOS)
						utils.OpenURIWithDefaultApplication(theURL)
					},
				},
				{
					Label:   "Configuration",
					Tooltip: "Information about configuring dev",
					Handler: func() {
						theURL := constants.RepositoryURLConfiguration
						log.Infof("opening '%s' for the '%s' platform", theURL, runtime.GOOS)
						utils.OpenURIWithDefaultApplication(theURL)
					},
				},
				{
					Type: systemtray.TypeSeparator,
				},
				{
					Label:   "Exit",
					Tooltip: "Shut down the Dev client tool",
					Handler: func() {
						log.Info("exit was clicked")
						systray.Quit()
					},
				},
			}, systemTrayStopped)
		},
	}
	return &cmd
}
