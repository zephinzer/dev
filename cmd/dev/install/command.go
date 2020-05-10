package install

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/install/repository"
	"github.com/usvc/dev/internal/constants"
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
