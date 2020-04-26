package client

import (
	"log"
	"runtime"

	"github.com/getlantern/systray"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/pkg/utils"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.ClientCanonicalNoun,
		Aliases: constants.ClientAliases,
		Short:   "starts the dev client as a background process to provide notifications",
		Run: func(command *cobra.Command, _ []string) {
			log.Println("adding system tray icon...")
			var (
				about *systray.MenuItem
				exit  *systray.MenuItem
			)
			systray.Run(func() {
				log.Println("initialising system tray...")
				systray.SetIcon(constants.SystrayIcon)
				systray.SetTooltip("Dev CLI tool")
				about = systray.AddMenuItem("About", "Display information about the Dev tool")
				systray.AddSeparator()
				exit = systray.AddMenuItem("Exit", "Shutdown the Dev tool")
				go func() {
					for {
						select {
						case <-about.ClickedCh:
							ourURL := "https://gitlab.com/usvc/utils/dev"
							log.Printf("opening '%s' for the '%s' platform", ourURL, runtime.GOOS)
							utils.OpenURIWithDefaultApplication(ourURL)
						case <-exit.ClickedCh:
							log.Println("exit was clicked")
							systray.Quit()
						}
					}
				}()
			}, func() {
				close(about.ClickedCh)
				close(exit.ClickedCh)
				log.Println("exiting system tray...")
			})
		},
	}
	return &cmd
}
