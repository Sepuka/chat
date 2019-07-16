package cmd

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sepuka/chat/src/def"
	"github.com/sepuka/chat/src/def/telegram"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(repeater)
}

var (
	repeater = &cobra.Command{
	Use: `repeater`,
	Example: `repeater -c /config/path`,
	Short: `repeater`,
	Long:  `repeater`,

	RunE: func(cmd *cobra.Command, args []string) error {
		var bot *tgbotapi.BotAPI
		if err := def.Container.Fill(telegram.TelegramDef, &bot); err != nil {
			return err
		}

		log.Printf("Authorized on account %s", bot.Self.UserName)

		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, _ := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
		return nil
	},
}
)