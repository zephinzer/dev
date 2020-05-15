package install

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/install/repository"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.InstallCanonicalVerb,
		Aliases: constants.InstallAliases,
		Short:   "sets stuff up on your machine",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(repository.GetCommand())
	return &cmd
}
