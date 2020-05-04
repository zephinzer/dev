package systemtray

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/log"
	systray "github.com/usvc/dev/internal/systemtray"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "system-tray",
		Aliases: []string{"systray"},
		Short:   "Tests the notifications",
		Run: func(command *cobra.Command, args []string) {
			stopped := make(chan struct{})
			systray.Start(systray.Menu{
				{
					Label:   "test item 1",
					Tooltip: "this is test item 1",
					Handler: func() {
						log.Info("item 1 was clicked!")
					},
				},
				{
					Type: systray.TypeSeparator,
				},
				{
					Label:   "test item 2",
					Tooltip: "this is test item 2",
					Handler: func() {
						log.Info("item 2 was clicked!")
					},
				},
				{
					Label:   "test item 3",
					Tooltip: "this is test item 3",
					Type:    systray.TypeFolder,
					Menu: &systray.Menu{
						{
							Label:   "nested item 1",
							Tooltip: "this is nested item 1",
							Handler: func() {
								log.Info("nested item 1 was clicked")
							},
						},
					},
				},
			}, stopped)
			<-stopped
		},
	}
	return &cmd
}
