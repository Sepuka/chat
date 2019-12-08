package log

import (
	"bytes"
	"errors"
	"fmt"
	http2 "net/http"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/config"
	"github.com/sepuka/chat/internal/def"
	"github.com/sepuka/chat/internal/def/http"
)

const (
	LoggerSyncerDef = `logger.syncer.def`
	TelegramBotDef  = `telegram.bot.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: TelegramBotDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					client = ctx.Get(http.ClientDef).(*http2.Client)
				)

				client.Timeout = time.Second * time.Duration(cfg.Log.Telegram.Timeout)

				bot, err := tgbotapi.NewBotAPIWithClient(cfg.Telegram.Token, client)
				if err != nil {
					return nil, errors.New(fmt.Sprintf(`unable connect to telegram api: %s`, err))
				}

				bot.Debug = cfg.Telegram.Debug

				return bot, nil
			},
		})
	})

	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: LoggerSyncerDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					botApi = def.Container.Get(TelegramBotDef)
				)

				return NewTelegramSyncer(botApi.(*tgbotapi.BotAPI), cfg.Log.Telegram.Channel), nil
			},
		})
	})
}

type TelegramSyncer struct {
	bot     *tgbotapi.BotAPI
	channel string
}

func NewTelegramSyncer(bot *tgbotapi.BotAPI, channel string) TelegramSyncer {
	return TelegramSyncer{
		bot:     bot,
		channel: channel,
	}
}

func (s TelegramSyncer) Sync() error {
	return nil
}

func (s TelegramSyncer) Write(p []byte) (n int, err error) {
	channel := strings.TrimPrefix(s.channel, `@`)
	parts := bytes.SplitN(p, []byte("\t"), 3)
	msg := tgbotapi.NewMessageToChannel(fmt.Sprintf(`@%s`, channel), string(bytes.TrimSpace(parts[2])))
	res, err := s.bot.Send(msg)

	return len(res.Text), err
}
