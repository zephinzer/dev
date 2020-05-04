package telegram

import (
	"github.com/spf13/cobra"
	"github.com/usvc/dev/cmd/dev/initialise/telegram/notifications"
	"github.com/usvc/dev/internal/constants"
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
