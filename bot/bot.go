package tgbot

import (
	"github.com/Syfaro/telegram-bot-api"
	
	"log"
	"../parser"
	//"io/ioutil"
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

	file := parser.Create_xls()
	ff := parser.File{file}
	ff.SaveAs("data.xlsx")
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.Text == "/send" {
			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "shalom")
			//bot.Send(msg)
			responce := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, "data.xlsx")
			_, err := bot.Send(responce)
			log.Print(err)
		}
	}
}