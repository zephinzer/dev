package initialise

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/initialise/database"
	"github.com/zephinzer/dev/cmd/dev/initialise/github"
	"github.com/zephinzer/dev/cmd/dev/initialise/telegram"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.InitialiseCanonicalNoun,
		Aliases: constants.InitialiseAliases,
		Short:   "individually initialise components within this tool",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(database.GetCommand())
	cmd.AddCommand(github.GetCommand())
	cmd.AddCommand(telegram.GetCommand())
	return &cmd
}
