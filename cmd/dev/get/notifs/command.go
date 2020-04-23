package notifs

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get/notifs/gitlab"
	"github.com/usvc/dev/cmd/dev/get/notifs/pivotaltracker"
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
	cmd.AddCommand(pivotaltracker.GetCommand())
	cmd.AddCommand(gitlab.GetCommand())
	return &cmd
}
