package notifications

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/internal/notifications"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.NotificationsCanonicalNoun,
		Aliases: constants.NotificationsAliases,
		Short:   "Tests the notifications",
		Run: func(command *cobra.Command, args []string) {
			hostname, _ := os.Hostname()
			errors := false
			title := "Hello, Dev"
			message := fmt.Sprintf(
				"This is a test notification from the dev app at %s",
				hostname,
			)
			link := "More about `dev` at " + constants.CanonicalRepositoryURL

			log.Info("triggering desktop test notification...")
			notificationError := notifications.TriggerDesktop(title, message, "./assets/icon/512-light.png")
			if notificationError != nil {
				log.Errorf("failed to trigger desktop notification: %s", notificationError)
				errors = true
			} else {
				log.Info("desktop test notification seems to have been triggered successfully")
			}

			if len(config.Global.Dev.Client.Notifications.Telegram.Token) > 0 &&
				len(config.Global.Dev.Client.Notifications.Telegram.ID) > 0 {
				botToken := config.Global.Dev.Client.Notifications.Telegram.Token
				targetChatID := config.Global.Dev.Client.Notifications.Telegram.ID

				log.Info("triggering telegram test notification...")
				chatID, chatIDtoIntError := strconv.Atoi(targetChatID)
				if chatIDtoIntError != nil {
					log.Errorf("failed to convert chat id '%s' to an integer: %s", targetChatID, chatIDtoIntError)
				} else if notificationError = notifications.TriggerTelegram(botToken, int64(chatID), fmt.Sprintf("*%s*\n\n_%s_\n\n[%s](%s)", title, message, link, constants.CanonicalRepositoryURL)); notificationError != nil {
					log.Errorf("failed to trigger telegram notification: %s", notificationError)
				} else {
					log.Info("telegram test notification seems to have been triggered successfully")
				}
			}
			if errors {
				os.Exit(1)
			}
		},
	}
	return &cmd
}
