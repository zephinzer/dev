package work

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get/work/pivotaltracker"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "work",
		Aliases: []string{"tickets", "stories", "tasks"},
		Short:   "Retrieves your work",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(pivotaltracker.GetCommand())
	return &cmd
}
