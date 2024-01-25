package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	currentStep map[int]*StepData
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot:         bot,
		currentStep: make(map[int]*StepData),
	}
}

func (b *Bot) Start() error {
	log.Printf("Авторизован на учетной записи: %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {

			// Команды.
			if update.Message.IsCommand() {
				b.updateCommand(update.Message)
				continue
			}

			// Обычный текст.
			b.updateMessage(update.Message)
		}
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
