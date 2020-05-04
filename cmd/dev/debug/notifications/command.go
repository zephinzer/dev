package notifications

import (
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
			errors := false
			log.Info("triggering desktop test notification...")
			notificationError := notifications.TriggerDesktop("Hello, Dev", "This is a test notification from the dev app", "/home/z/gitlab.com/usvc/utils/dev/assets/icon/512-light.png")
			if notificationError != nil {
				log.Errorf("failed to trigger notification: %s", notificationError)
				errors = true
			} else {
				log.Info("desktop test notification seems to have been triggered successfully")
			}

			if len(config.Global.Dev.Client.Notifications.Telegram.Token) > 0 &&
				len(config.Global.Dev.Client.Notifications.Telegram.ID) > 0 {
				log.Info("triggering telegram test notification...")
				bot, newBotError := tgbotapi.NewBotAPI(config.Global.Dev.Client.Notifications.Telegram.Token)
				if newBotError != nil {
					log.Errorf("failed to create a new bot instance: %s", newBotError)
					os.Exit(1)
				} else {
					chatID, chatIDtoIntError := strconv.Atoi(config.Global.Dev.Client.Notifications.Telegram.ID)
					if chatIDtoIntError != nil {
						log.Errorf("failed to convert %s to an integer", config.Global.Dev.Client.Notifications.Telegram.ID)
					} else {
						bot.Send(tgbotapi.NewMessage(int64(chatID), "Hello, Dev. This is a test notification from the dev app"))
						log.Info("telegram test notification seems to have been triggered successfully")
					}
				}
			}
			if errors {
				os.Exit(1)
			}
		},
	}
	return &cmd
}
