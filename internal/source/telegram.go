package source

import (
	"fmt"

	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
	"go.uber.org/zap"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Telegram struct {
	commands command.HandlerMap
	bot      *tgbotapi.BotAPI
	logger   *zap.SugaredLogger
}

func NewTelegram(
	commandsMap command.HandlerMap,
	bot *tgbotapi.BotAPI,
	logger *zap.SugaredLogger,
) *Telegram {
	return &Telegram{
		commands: commandsMap,
		bot:      bot,
		logger:   logger,
	}
}

func (hosting *Telegram) Listen() error {
	hosting.logger.Infof(`authorized on account "%s"`, hosting.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, _ := hosting.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		req := context.NewRequest(update.Message.From.UserName, domain.Telegram, update.Message.Text)
		hosting.logger.
			Info(
				`got telegram command`,
				zap.String(`user`, req.GetLogin()),
				zap.String(`command`, req.GetCommand()),
				zap.Strings(`args`, req.GetArgs()),
			)

		var result *command.Result
		var err error
		if f, ok := hosting.commands[update.Message.Text]; ok {
			hosting.sendAnswer(update.Message, fmt.Sprintf(`command '%s' accepted`, update.Message.Text))

			result, err = f.Exec(req)
			var answer string
			if err != nil {
				answer = fmt.Sprintf(`error: %s`, err.Error())
			} else {
				answer = fmt.Sprintf(`command result: %s`, result.Response)
			}
			hosting.sendAnswer(update.Message, answer)
		} else {
			hosting.sendAnswer(update.Message, `unknown command`)
		}
	}

	return nil
}

func (hosting *Telegram) sendAnswer(srcMsg *tgbotapi.Message, text string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(srcMsg.Chat.ID, text)
	msg.ReplyToMessageID = srcMsg.MessageID
	return hosting.bot.Send(msg)
}
