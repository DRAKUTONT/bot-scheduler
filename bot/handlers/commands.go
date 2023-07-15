package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (handler *Handler) StartAndHelp(update tgbotapi.Update) error {
	result, err := handler.database.IsUserExists(int(update.Message.From.ID))
	if err != nil {
		return err
	}

	if !result {
		handler.database.AddUser(int(update.Message.From.ID))
	}

	message := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Вот список моих команд:\n/help - список команд\n/stop - остановка бота\n/activate - запуск бота\n")

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

func (handler *Handler) Stop(update tgbotapi.Update) error {
	if err := handler.database.UpdateIsActive(int(update.Message.From.ID), false); err != nil {
		return err
	}

	message := tgbotapi.NewMessage(update.Message.Chat.ID, "Бот остановлен, чтобы продолжить работу бота напишите /activate")

	if err := handler.SendMessage(message); err != nil {
		return err
	}

	return nil
}

func (handler *Handler) Activate(update tgbotapi.Update) error {
	if err := handler.database.UpdateIsActive(int(update.Message.From.ID), true); err != nil {
		return err
	}

	message := tgbotapi.NewMessage(update.Message.Chat.ID, "Бот продолжил работу, чтобы остановить работу бота напишите /stop")

	if err := handler.SendMessage(message); err != nil {
		return err
	}

	return nil
}
