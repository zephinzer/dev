package initialise

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/initialise/database"
	"github.com/usvc/dev/cmd/dev/initialise/pivotaltracker"
	"github.com/usvc/dev/cmd/dev/initialise/telegram"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.InitialiseCanonicalNoun,
		Aliases: constants.InitialiseAliases,
		Short:   "Initialises components within the dev CLI tool",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(database.GetCommand())
	cmd.AddCommand(pivotaltracker.GetCommand())
	cmd.AddCommand(telegram.GetCommand())
	return &cmd
}
