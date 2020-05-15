package github

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/get/github/account"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.GithubCanonicalNoun,
		Aliases: constants.GithubAliases,
		Short:   "Retrieves information from Github",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(account.GetCommand())
	return &cmd
}
