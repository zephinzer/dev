package client

import (
	"runtime"

	"github.com/getlantern/systray"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/internal/systemtray"
	"github.com/usvc/dev/pkg/utils"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.ClientCanonicalNoun,
		Aliases: constants.ClientAliases,
		Short:   "starts the dev client as a background process to provide notifications",
		Run: func(command *cobra.Command, _ []string) {
			log.Debug("adding system tray icon...")
			systemTrayStopped := make(chan struct{})
			systemtray.Start(systemtray.Menu{
				{
					Label:   "About",
					Tooltip: "Display information about the Dev client tool",
					Handler: func() {
						ourURL := "https://gitlab.com/usvc/utils/dev"
						log.Infof("opening '%s' for the '%s' platform", ourURL, runtime.GOOS)
						utils.OpenURIWithDefaultApplication(ourURL)
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
