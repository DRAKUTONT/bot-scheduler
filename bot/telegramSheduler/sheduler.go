package telegramSheduler

import (
	"antiProcrastinationBotModule/database"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Shedule(bot *tgbotapi.BotAPI, database *database.Database) error {
	fmt.Println("Sheduler is active")
	users, err := database.GetUsers()
	if err != nil {
		return fmt.Errorf("can't get user(shedule func): %w", err)
	}
	for _, user := range users {
		message := tgbotapi.NewMessage(int64(user.UserID), "Ваш список задач на сегодня:"+user.TaskList)

		if _, err := bot.Send(message); err != nil {
			return fmt.Errorf("can't send message: %w", err)
		}
	}
	return nil
}
