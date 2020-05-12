package notifications

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func TriggerTelegram(botToken string, chatID int64, message string) error {
	bot, newBotError := tgbotapi.NewBotAPI(botToken)
	if newBotError != nil {
		return fmt.Errorf("failed to create a new bot instance: %s", newBotError)
	}
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = tgbotapi.ModeMarkdown
	if _, sendMessageError := bot.Send(msg); sendMessageError != nil {
		return fmt.Errorf("failed to send the message: %s", sendMessageError)
	}
	return nil
}
