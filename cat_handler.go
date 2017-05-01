package main

import (
	"github.com/zhulik/margelet"
	"gopkg.in/telegram-bot-api.v4"
	"math/rand"
	"time"
)

const catURL = "http://thecatapi.com/api/images/get?format=src&type=jpg"

type CatHandler struct {
}

func (h CatHandler) HandleCommand(message margelet.Message) error {
	return sendCat(message.Message().Chat.ID, message.Bot(), message.Message().MessageID)
}

func (h CatHandler) HelpMessage() string {
	return "Send image with cat"
}

func downloadCat() ([]byte, error) {
	return downloadFromUrl(catURL)
}

func sendCat(chatID int64, bot margelet.MargeletAPI, replyTo ...int) error {
	bytes, err := downloadCat()
	if err != nil {
		return err
	}

	bot.Send(tgbotapi.NewChatAction(chatID, "upload_photo"))
	msg := tgbotapi.NewPhotoUpload(chatID,
		tgbotapi.FileBytes{Name: "cat.jpg", Bytes: bytes})
	if len(replyTo) > 0 {
		msg.ReplyToMessageID = replyTo[0]
	}
	bot.Send(msg)
	return nil
}

func randomCatSender(bot *margelet.Margelet) {
	for {
		for _, chatID := range bot.ChatRepository.All() {
			if bot.ChatConfigRepository.Get(chatID) == "yes" {
				go sendCat(chatID, bot)
			}
		}
		time.Sleep(time.Duration(rand.Intn(59)+1) * time.Second)
	}
}
