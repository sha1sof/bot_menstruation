package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b Bot) updateCommand(message *tgbotapi.Message) {
	log.Printf("[%s]:- Command: %s", message.From.UserName, message.Text)

	switch message.Command() {
	case start:
		msg := tgbotapi.NewMessage(message.Chat.ID, startMessage)
		b.bot.Send(msg)
	case help:
		msg := tgbotapi.NewMessage(message.Chat.ID, helpMessage)
		b.bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, unknownCommand)
		b.bot.Send(msg)
	}
}

func (b Bot) updateMessage(message *tgbotapi.Message) {
	log.Printf("[%s]:- %s", message.From.UserName, message.Text)

}
