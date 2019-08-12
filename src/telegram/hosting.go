package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sepuka/chat/src/def"
	"github.com/sepuka/chat/src/def/telegram"
	"github.com/sepuka/chat/src/domain"
)

func Hosting() error {
	var (
		bot *tgbotapi.BotAPI
		msg tgbotapi.MessageConfig
	)
	if err := def.Container.Fill(telegram.BotDef, &bot); err != nil {
		return err
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		cmd, err := domain.NewCommand(update.Message.Text)
		if err != nil {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(`command '%s' accepted`, cmd))
		}

		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

	return nil
}
