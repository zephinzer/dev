package work

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get/work/pivotaltracker"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.WorkCanonicalNoun,
		Aliases: constants.WorkAliases,
		Short:   "Retrieves your work",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(pivotaltracker.GetCommand())
	return &cmd
}
