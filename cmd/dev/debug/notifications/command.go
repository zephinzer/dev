package notifications

import (
	"github.com/0xAX/notificator"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/constants"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.NotificationsCanonicalNoun,
		Aliases: constants.NotificationsAliases,
		Short:   "Tests the notifications",
		Run: func(command *cobra.Command, args []string) {
			notificator.New(notificator.Options{
				DefaultIcon: "https://image.flaticon.com/icons/png/512/119/119060.png",
				AppName:     "dev",
			}).Push("Hello, Dev", "This is a test notification from the dev app", "https://image.flaticon.com/icons/png/512/119/119060.png", notificator.UR_NORMAL)
		},
	}
	return &cmd
}
