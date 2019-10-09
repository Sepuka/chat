package source

import (
	"chat/src/command"
	"chat/src/context"
	"chat/src/domain"
	"fmt"
	"go.uber.org/zap"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Telegram struct {
	commands map[string]command.Executor
	bot      *tgbotapi.BotAPI
	logger   *zap.SugaredLogger
}

func NewTelegram(
	commandsMap map[string]command.Executor,
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
	var (
		msg tgbotapi.MessageConfig
	)

	log.Printf("Authorized on account %s", hosting.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, _ := hosting.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		req := context.NewRequest(update.Message.From.UserName, update.Message.Text)
		hosting.logger.
			Info(
				`got telegram command`,
				zap.String(`user`, req.GetLogin()),
				zap.String(`command`, req.GetCommand()),
				zap.Strings(`args`, req.GetArgs()),
			)

		if f, ok := hosting.commands[update.Message.Text]; ok {
			f.Exec(req)
		}

		cmd, err := domain.NewCommand(update.Message.Text)
		if err != nil {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(`command '%s' accepted`, cmd))
		}

		msg.ReplyToMessageID = update.Message.MessageID

		hosting.bot.Send(msg)
	}

	return nil
}
