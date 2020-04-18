package account

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get/account/pivotaltracker"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "account",
		Aliases: []string{"acc", "a"},
		Short:   "Retrieves account information",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(pivotaltracker.GetCommand())
	return &cmd
}
