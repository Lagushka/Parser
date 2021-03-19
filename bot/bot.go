package tgbot

import (
	"github.com/Syfaro/telegram-bot-api"
	
	"log"
	"../parser"
	//"io/ioutil"
)

type tBot struct {
	bot tgbotapi.BotAPI
}

var messages []tgbotapi.Message

func sendMessage(msg tgbotapi.Message) {

}

func Work() {
	bot, err := tgbotapi.NewBotAPI("1669961010:AAGWt0qbTW_HwyraRnN1Q3H6j62GqeT1jPs")

	tgbot := tBot { bot }
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
	sendButton := tgbotapi.NewKeyboardButton("Send Chart")
	row := tgbotapi.NewKeyboardButtonRow(sendButton)
	keyboard:= tgbotapi.NewReplyKeyboard(row)

	for update := range updates {

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "shalom")

		if update.Message.Text == "/send" {
			msg.ReplyMarkup = keyboard
			_, err = bot.Send(msg)
			log.Print(err)
		}

		if update.Message.Text == "Send Chart" {
			responce := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, "data.xlsx")
			_, err = bot.Send(responce)
			log.Print(err)

			rmKeyboard := tgbotapi.NewRemoveKeyboard(true)
			msg.ReplyMarkup = rmKeyboard
			_, err := bot.Send(msg)
			log.Print(err)
		}
	}
}