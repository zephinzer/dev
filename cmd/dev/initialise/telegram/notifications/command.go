package notifications

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.NotificationsCanonicalNoun,
		Aliases: constants.NotificationsAliases,
		Short:   "Initialises telegram notifications",
		Run: func(command *cobra.Command, args []string) {
			if len(config.Global.Dev.Client.Notifications.Telegram.Token) == 0 {
				log.Error("you need a bot token before initialising the notifications for telegram")
				log.Error("talk to the @BotFather at https://t.me/BotFather and paste your bot's token into your YAML configuration at `dev.client.notifications.telegram.token`")
				os.Exit(1)
			}
			log.Debug("creating telegram bot controller to retrieve your user id...")
			bot, newBotError := tgbotapi.NewBotAPI(config.Global.Dev.Client.Notifications.Telegram.Token)
			if newBotError != nil {
				log.Errorf("failed to create a new bot instance: %s", newBotError)
				os.Exit(1)
			}
			botInfo, getBotInfoError := bot.GetMe()
			if getBotInfoError != nil {
				log.Errorf("failed to retrieve bot information using provided token: %s", getBotInfoError)
				os.Exit(1)
			}
			log.Debugf("using bot '%s' (@%s)", botInfo.String(), botInfo.UserName)
			log.Debug("setting up update configurations...")
			updateConfig := tgbotapi.NewUpdate(0)
			updateConfig.Timeout = 10
			updatesStream, updatesStreamError := bot.GetUpdatesChan(updateConfig)
			if updatesStreamError != nil {
				log.Errorf("failed to retrieve update stream for bot: %s", updatesStreamError)
				os.Exit(1)
			}
			log.Debugf("starting listener for telegram bot @%s...", botInfo.UserName)
			log.Debugf("go to https://t.me/%s and click on the **Start** button to retrieve your telegram id", botInfo.UserName)
			for update := range updatesStream {
				if update.Message == nil {
					continue
				}
				log.Infof("your telegram name is    : %s %s", update.Message.From.FirstName, update.Message.From.LastName)
				log.Infof("your telegram username is: %s", update.Message.From.UserName)
				log.Infof("your telegram **id** is  : %v", update.Message.From.ID)
				log.Printf("if your above telegram info is correct, insert/merge the following YAML configuration with your own:\n")
				log.Printf("dev:\n")
				log.Printf("  client:\n")
				log.Printf("    notifications:\n")
				log.Printf("      telegram:\n")
				log.Printf("        token: ...\n")
				log.Printf("        id: \"%v\"\n", update.Message.From.ID)
				bot.StopReceivingUpdates()
				break
			}
		},
	}
	return &cmd
}
