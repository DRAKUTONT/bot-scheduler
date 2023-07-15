package handlers

import (
	"fmt"
	"strings"

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

	message := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Вот список моих команд:\n/help - список команд\n/stop - остановка бота\n/activate - запуск бота\n/create_task <список заданий> - создать список заданий\n/my_task - текущие задачи")

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

func (handler *Handler) CreateTask(update tgbotapi.Update) error {
	_, task, _ := strings.Cut(update.Message.Text, " ")
	fmt.Println(task)
	message := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if task == "" {
		message.Text = "Вы не указали задание"
	} else {
		message.Text = "Задача успешно создана"
		handler.database.UpdateTask(int(update.Message.From.ID), task)
	}

	if err := handler.SendMessage(message); err != nil {
		return err
	}

	return nil
}

func (handler *Handler) GetTask(update tgbotapi.Update) error {
	task, err := handler.database.GetTask(int(update.Message.From.ID))

	if err != nil {
		return fmt.Errorf("can't get task: %w", err)
	}

	message := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if task == "" {
		message.Text = "У вас нет текущих заданий, чтобы добавить задание введите /create_task <список заданий>"
	} else {
		message.Text = "Ваши задачи:\n" + task
	}

	if err := handler.SendMessage(message); err != nil {
		return err
	}

	return nil
}
