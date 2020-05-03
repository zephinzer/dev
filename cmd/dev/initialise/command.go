package initialise

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/initialise/database"
	"github.com/usvc/dev/cmd/dev/initialise/pivotaltracker"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.InitialiseCanonicalNoun,
		Aliases: constants.InitialiseAliases,
		Short:   "Initialises the dev CLI tool for persistent use",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(database.GetCommand())
	cmd.AddCommand(pivotaltracker.GetCommand())
	return &cmd
}
