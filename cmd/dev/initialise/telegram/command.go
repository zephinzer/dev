package telegram

import (
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/initialise/telegram/notifications"
	"github.com/zephinzer/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.TelegramCanonicalNoun,
		Aliases: constants.TelegramAliases,
		Short:   "Initialises telegram related stuff",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(notifications.GetCommand())
	return &cmd
}
