package account

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/get/account/github"
	"github.com/usvc/dev/cmd/dev/get/account/gitlab"
	"github.com/usvc/dev/cmd/dev/get/account/pivotaltracker"
	"github.com/usvc/dev/cmd/dev/get/account/trello"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.AccountCanonicalNoun,
		Aliases: constants.AccountAliases,
		Short:   "Retrieves account information",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(gitlab.GetCommand())
	cmd.AddCommand(github.GetCommand())
	cmd.AddCommand(pivotaltracker.GetCommand())
	cmd.AddCommand(trello.GetCommand())
	return &cmd
}
