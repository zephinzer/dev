package notifications

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get/notifications/gitlab"
	"github.com/usvc/dev/cmd/dev/get/notifications/pivotaltracker"
	"github.com/usvc/dev/cmd/dev/get/notifications/trello"
	"github.com/usvc/dev/internal/constants"
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
