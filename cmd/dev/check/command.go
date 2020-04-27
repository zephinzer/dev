package check

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/check/software"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.CheckCanonicalVerb,
		Aliases: constants.CheckAliases,
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(software.GetCommand())
	return &cmd
}
