package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	client := &http.Client{}
	proxy, err := url.Parse("http://3.212.104.192:3128")
	if err != nil {
		log.Panic(err)
	}

	client.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}
	bot, err := tgbotapi.NewBotAPIWithClient("", client)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
