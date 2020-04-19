package initialise

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/initialise/database"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "initialise",
		Aliases: []string{"initialize", "init", "i"},
		Short:   "Initialises the dev CLI tool for persistent use",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(database.GetCommand())
	return &cmd
}
