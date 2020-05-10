package repository

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.RepositoryCanonicalVerb,
		Aliases: constants.RepositoryAliases,
		Short:   "sets up repositories as defined in the configuration",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	return &cmd
}
