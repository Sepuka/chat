package source

import (
	"fmt"
	"strings"

	"github.com/sepuka/chat/internal/middleware"

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
	handler  middleware.HandlerFunc
}

func NewTelegram(
	commandsMap command.HandlerMap,
	bot *tgbotapi.BotAPI,
	logger *zap.SugaredLogger,
	handler middleware.HandlerFunc,
) *Telegram {
	return &Telegram{
		commands: commandsMap,
		bot:      bot,
		logger:   logger,
		handler:  handler,
	}
}

func (hosting *Telegram) Listen() error {
	var (
		splitedCmd []string
		req        *context.Request
	)
	hosting.logger.Infof(`authorized on account "%s"`, hosting.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, _ := hosting.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		splitedCmd = strings.Split(update.Message.Text, ` `)
		switch len(splitedCmd) {
		case 0, 1:
			req = context.NewRequest(update.Message.From.UserName, domain.Telegram, update.Message.Text)
		default:
			req = context.NewRequest(update.Message.From.UserName, domain.Telegram, splitedCmd[0], splitedCmd[1:]...)
		}

		hosting.logger.
			Info(
				`got telegram command`,
				zap.String(`user`, req.GetLogin()),
				zap.String(`command`, req.GetCommand()),
				zap.Strings(`args`, req.GetArgs()),
			)

		if finalHandler, ok := hosting.commands[req.GetCommand()]; ok {
			hosting.sendAnswer(update.Message, fmt.Sprintf(`command '%s' accepted`, req.GetCommand()))

			var (
				resp   = &command.Result{}
				err    error
				answer string
			)
			err = hosting.handler(finalHandler, req, resp)

			if err != nil {
				answer = fmt.Sprintf(`error: %s`, err.Error())
			} else {
				answer = string(resp.Response)
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
