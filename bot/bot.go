package tgbot

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	//"../parser"
)

func Work() {
	bot, err := tgbotapi.NewBotAPI("1669961010:AAGWt0qbTW_HwyraRnN1Q3H6j62GqeT1jPs")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on this account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//file := parser.Create_xls()

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		s := update.Message.Command()
		log.Printf(s)

		if s == "send" {
			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "shalom")
			//bot.Send(msg)
			responce := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, "../tree.html")
			bot.Send(responce)
		}
	}
}