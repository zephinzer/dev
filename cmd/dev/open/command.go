package open

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/open/repo"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.OpenCanonicalVerb,
		Aliases: constants.OpenAliases,
		Short:   "convenience sub-commands to open stuff",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(repo.GetCommand())
	return &cmd
}
