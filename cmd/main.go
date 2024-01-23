package main

import (
	"log"
	"os"

	"github.com/sha1sof/bot_tg_menstruation.git/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./config/.env")
	keyAPI := os.Getenv("BOT_KEY")

	bot, err := tgbotapi.NewBotAPI(keyAPI)
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
