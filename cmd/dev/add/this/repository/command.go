package repository

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.RepositoryCanonicalNoun,
		Aliases: constants.RepositoryAliases,
		Short:   "adds the current repository you're in to dev's configuration",
		Run:     run,
	}
	return &cmd
}

func run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
