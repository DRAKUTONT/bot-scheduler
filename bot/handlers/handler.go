package handlers

import (
	"antiProcrastinationBotModule/database"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TODO: прикрутить gocron в хэндлер
type Handler struct {
	bot     *tgbotapi.BotAPI
	database database.Database
	timeout int
}

func New(bot *tgbotapi.BotAPI, database *database.Database, timeout int) *Handler {
	return &Handler{
		bot:     bot,
		database: *database,
		timeout: timeout,
	}
}

func (handler *Handler) HandleMessage() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := handler.bot.GetUpdatesChan(u)

	for update := range updates {
		if !handler.isValidMessage(&update) {
			continue
		}

		switch update.Message.Command() {
		case "start", "help":
			handler.StartAndHelp(update)
			fmt.Println(update.Message.Entities)
		default:
			handler.UnknownMessage(update)
		}

	}
	return nil
}

func (handler *Handler) SendMessage(message tgbotapi.MessageConfig) error {
	if _, err := handler.bot.Send(message); err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}

	return nil
}

func (handler *Handler) isValidMessage(update *tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}

	return true
}
