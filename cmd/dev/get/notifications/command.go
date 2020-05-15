package notifications

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/get/notifications/gitlab"
	"github.com/zephinzer/dev/cmd/dev/get/notifications/pivotaltracker"
	"github.com/zephinzer/dev/cmd/dev/get/notifications/trello"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.NotificationsCanonicalNoun,
		Aliases: constants.NotificationsAliases,
		Short:   "Retrieves notifications",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(gitlab.GetCommand())
	cmd.AddCommand(pivotaltracker.GetCommand())
	cmd.AddCommand(trello.GetCommand())
	return &cmd
}
