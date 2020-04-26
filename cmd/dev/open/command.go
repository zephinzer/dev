package open

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/open/repo"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.OpenCanonicalVerb,
		Aliases: constants.OpenAliases,
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(repo.GetCommand())
	return &cmd
}
