package gitlab

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/get/gitlab/account"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GitlabCanonicalNoun,
		Aliases: constants.GitlabAliases,
		Short:   "Retrieves information from Gitlab",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(account.GetCommand())
	return &cmd
}
