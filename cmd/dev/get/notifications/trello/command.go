package trello

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/trello"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "trello",
		Aliases: []string{"tr"},
		Short:   "Retrieves notifications from Trello",
		Run: func(command *cobra.Command, args []string) {
			notifications, err := trello.GetNotifications(config.Global.Platforms.Trello.AccessKey, config.Global.Platforms.Trello.AccessToken)
			if err != nil {
				panic(err)
			}
			log.Printf("# notifications from trello\n\n%s", notifications.String())
			log.Printf("> you have a total of %v unread notifications on your linked trello account", len(*notifications))
		},
	}
	return &cmd
}
