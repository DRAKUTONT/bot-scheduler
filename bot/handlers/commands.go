package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (handler *Handler) StartAndHelp(update tgbotapi.Update) error {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, "start")

	if err := handler.SendMessage(message); err != nil {
		return err
	}

	return nil
}

func (handler *Handler) UnknownMessage(update tgbotapi.Update) error {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, "Я тебя не понимаю. Напиши /help, чтобы ознакомиться co списком команд")

	if err := handler.SendMessage(message); err != nil {
		return err
	}

	return nil
}
