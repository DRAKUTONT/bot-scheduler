package main

import (
	"antiProcrastinationBotModule/bot/handlers"
	"antiProcrastinationBotModule/bot/telegramSheduler"
	"antiProcrastinationBotModule/database"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5109491914:AAGGj3Vm6_Lj16KGD6e-AsocVIp_USuF3jk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	database, err := database.New("database/database.db")
	if err != nil {
		log.Panic(err)
	}

	database.Init()

	sheduler := gocron.NewScheduler(time.Now().Location())
	sheduler.Every(1).Day().At("10:00;13:00;19:00").Do(telegramSheduler.Shedule, bot, database)
	sheduler.StartAsync()

	handler := handlers.New(bot, database, 60)
	handler.HandleMessage()
}
