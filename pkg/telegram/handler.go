package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StepData struct {
	Step int
}

const (
	StepStart = iota + 1
	StepManNik
	StepManConfirmation
	StepManWait

	StepWomanData
	StepWomanCorrection
	StepWomanNik
	StepWomanConfirmation
)

func (b Bot) updateCommand(message *tgbotapi.Message) {
	log.Printf("[%s]:- Command: %s", message.From.UserName, message.Text)

	switch message.Command() {
	case start:
		b.currentStep[int(message.From.ID)] = &StepData{Step: StepStart}
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
	if message.Text == "Мужчина" || message.Text == "мужчина" || message.Text == "М" || message.Text == "м" {
		b.currentStep[int(message.From.ID)] = &StepData{Step: StepManNik}
		msg := tgbotapi.NewMessage(message.Chat.ID, stepManNik)
		b.bot.Send(msg)
	} else if message.Text == "Женщина" || message.Text == "женщина" || message.Text == "Ж" || message.Text == "ж" {
		b.currentStep[int(message.From.ID)] = &StepData{Step: StepWomanData}
		msg := tgbotapi.NewMessage(message.Chat.ID, stepWomanData)
		b.bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, unknownGender)
		b.bot.Send(msg)
	}
}

func (b *Bot) handleStepNikMan(message *tgbotapi.Message) {
	//example := message.Text
	//example = strings.ReplaceAll(example, "@", "")
	msg := tgbotapi.NewMessage(message.Chat.ID, stepManConfirmation)
	b.bot.Send(msg)
	b.currentStep[int(message.From.ID)] = &StepData{Step: StepManConfirmation}
}

func (b *Bot) handleStepWomanData(message *tgbotapi.Message) {
	//data:=message.Text
	msg := tgbotapi.NewMessage(message.Chat.ID, stepWomanCorrection)
	b.bot.Send(msg)
	b.currentStep[int(message.From.ID)] = &StepData{Step: StepWomanCorrection}
}

func (b *Bot) handleStepWomanCorrection(message *tgbotapi.Message) {
	//day:=message.Text
	msg := tgbotapi.NewMessage(message.Chat.ID, stepWomanNik)
	b.bot.Send(msg)
	b.currentStep[int(message.From.ID)] = &StepData{Step: StepWomanNik}
}

func (b *Bot) handleStepWomanNik(message *tgbotapi.Message) {
	//example := message.Text
	//example = strings.ReplaceAll(example, "@", "")
	msg := tgbotapi.NewMessage(message.Chat.ID, stepWomanConfirmation)
	b.bot.Send(msg)
	b.currentStep[int(message.From.ID)] = &StepData{Step: StepWomanConfirmation}
}
