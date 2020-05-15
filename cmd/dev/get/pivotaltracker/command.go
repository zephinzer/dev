package pivotaltracker

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/get/pivotaltracker/account"
	"github.com/zephinzer/dev/cmd/dev/get/pivotaltracker/notifications"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.PivotalTrackerCanonicalNoun,
		Aliases: constants.PivotalTrackerAliases,
		Short:   "Retrieves information from PivotalTracker",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(account.GetCommand())
	cmd.AddCommand(notifications.GetCommand())
	return &cmd
}
