package telegram

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sha1sof/bot_menstruation.git/pkg/telegram/db"
)

type StepData struct {
	Step int
}

var partnerMan, partnerWoman string
var averageDuration uint
var startMenstruation time.Time

var startMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Мужчина"),
		tgbotapi.NewKeyboardButton("Женщина"),
	),
)

func (b Bot) updateCommand(message *tgbotapi.Message) {
	log.Printf("[%s]:- Command: %s", message.From.UserName, message.Text)

	switch message.Command() {
	case startCommand:
		b.currentStep[int(message.From.ID)] = &StepData{Step: StepStart}
		msg := tgbotapi.NewMessage(message.Chat.ID, startMessage)
		msg.ReplyMarkup = startMenu
		b.bot.Send(msg)
	case helpCommand:
		msg := tgbotapi.NewMessage(message.Chat.ID, helpMessage)
		b.bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, unknownCommand)
		b.bot.Send(msg)
	}
}

func (b Bot) updateMessage(message *tgbotapi.Message) {
	log.Printf("[%s]:- %s", message.From.UserName, message.Text)

	stepData, exists := b.currentStep[int(message.From.ID)]

	switch {
	case exists && stepData.Step == StepStart:
		b.handleStepAskGender(message)

	case exists && stepData.Step == StepManNik:
		b.handleStepNikMan(message)

	case exists && stepData.Step == StepWomanData:
		b.handleStepWomanData(message)

	case exists && stepData.Step == StepWomanCorrection:
		b.handleStepWomanCorrection(message)

	case exists && stepData.Step == StepWomanNik:
		b.handleStepWomanNik(message)
	}
}

func (b *Bot) handleStepAskGender(message *tgbotapi.Message) {
	if message.Text == "Мужчина" || message.Text == "мужчина" {
		b.currentStep[int(message.From.ID)] = &StepData{Step: StepManNik}
		msg := tgbotapi.NewMessage(message.Chat.ID, stepManNik)
		msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{RemoveKeyboard: true}
		b.bot.Send(msg)
	} else if message.Text == "Женщина" || message.Text == "женщина" {
		b.currentStep[int(message.From.ID)] = &StepData{Step: StepWomanData}
		msg := tgbotapi.NewMessage(message.Chat.ID, stepWomanData)
		msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{RemoveKeyboard: true}
		b.bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, unknownGender)
		b.bot.Send(msg)
	}
}

func (b *Bot) handleStepNikMan(message *tgbotapi.Message) {
	partnerMan, isValid := checkPartnerInvalid(message.Text)
	if !isValid {
		msg := tgbotapi.NewMessage(message.Chat.ID, unknownNik)
		b.bot.Send(msg)
		return
	}

	bo, err := db.AddMan(uint(message.Chat.ID), message.From.UserName, partnerMan)
	if err != nil {
		log.Print("Error: ", err)
		return
	}
	if bo {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Ваша девушка уже есть в базе данных")
		b.bot.Send(msg)
		b.currentStep[int(message.From.ID)] = &StepData{Step: StepManConfirmation}
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, stepManConfirmation)
	b.bot.Send(msg)
	b.currentStep[int(message.From.ID)] = &StepData{Step: StepManConfirmation}
}

func (b *Bot) handleStepWomanData(message *tgbotapi.Message) {
	var err error
	startMenstruation, err = time.Parse("02.01.2006", message.Text)

	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Ошибка формата даты\nПопробуйсте снова!")
		b.bot.Send(msg)
		return
	}

	fmt.Println("Parsed startMenstruation:", startMenstruation)

	msg := tgbotapi.NewMessage(message.Chat.ID, stepWomanCorrection)
	b.bot.Send(msg)
	b.currentStep[int(message.From.ID)] = &StepData{Step: StepWomanCorrection}
}
func (b *Bot) handleStepWomanCorrection(message *tgbotapi.Message) {
	text, err := strconv.Atoi(message.Text)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Попробуйте ввести снова\nТолько число дней")
		b.bot.Send(msg)
		log.Print("Error: Convert to uint: ", err)
	}

	averageDuration = uint(text)
	msg := tgbotapi.NewMessage(message.Chat.ID, stepWomanNik)
	b.bot.Send(msg)
	b.currentStep[int(message.From.ID)] = &StepData{Step: StepWomanNik}
}

func (b *Bot) handleStepWomanNik(message *tgbotapi.Message) {
	partnerWoman, isValid := checkPartnerInvalid(message.Text)
	if !isValid {
		msg := tgbotapi.NewMessage(message.Chat.ID, unknownNik)
		b.bot.Send(msg)
		return
	}
	bo, err := db.AddWoman(uint(message.Chat.ID), message.From.UserName, partnerWoman, averageDuration, startMenstruation)
	if err != nil {
		log.Print("Error: ", err)
		return
	}
	if bo {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Ваша парень уже есть в базе данных")
		b.bot.Send(msg)
		b.currentStep[int(message.From.ID)] = &StepData{Step: StepManConfirmation}
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, stepWomanConfirmation)
	b.bot.Send(msg)
	b.currentStep[int(message.From.ID)] = &StepData{Step: StepWomanConfirmation}
}

func checkPartnerInvalid(userName string) (string, bool) {
	userName = strings.ReplaceAll(userName, "@", "")

	for _, r := range userName {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') && (r != '_' && r != '@') {
			return "", false
		}
	}
	return userName, true
}
