package pivotaltracker

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/initialise/pivotaltracker/database"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.PivotalTrackerCanonicalNoun,
		Aliases: constants.PivotalTrackerAliases,
		Short:   "Initialises pivotal tracker related stuff",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(database.GetCommand())
	return &cmd
}