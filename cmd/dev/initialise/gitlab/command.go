package gitlab

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/initialise/gitlab/database"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GitlabCanonicalNoun,
		Aliases: constants.GitlabAliases,
		Short:   "Initialises gitlab related stuff",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(database.GetCommand())
	return &cmd
}
