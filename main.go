package main

import (
	"antiProcrastinationBotModule/bot/handlers"
	"antiProcrastinationBotModule/database"
	"log"

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
	if err != nil{
		log.Panic(err)
	}

	handler := handlers.New(bot, 60)
	handler.HandleMessage()
}
