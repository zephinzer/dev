package trello

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/get/trello/account"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.TrelloCanonicalNoun,
		Aliases: constants.TrelloAliases,
		Short:   "Retrieves information from Trello",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(account.GetCommand())
	return &cmd
}
