package notifs

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get/notifs/pivotaltracker"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "notifs",
		Short: "Retrieves notifications",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(pivotaltracker.GetCommand())
	return &cmd
}
