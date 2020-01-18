package source

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sepuka/chat/internal/def/middleware"
	middleware2 "github.com/sepuka/chat/internal/middleware"

	commandDef "github.com/sepuka/chat/internal/def/command"

	"github.com/sepuka/chat/internal/config"

	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/def"
	httpClient "github.com/sepuka/chat/internal/def/http"
	"github.com/sepuka/chat/internal/def/log"
	"github.com/sepuka/chat/internal/source"
	"go.uber.org/zap"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	TelegramDef = `command.handler.telegram.def`
	BotDef      = `command.source.telegram.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: BotDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					client = ctx.Get(httpClient.ClientDef).(*http.Client)
				)

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
			Name: TelegramDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					logger     = def.Container.Get(log.LoggerDef).(*zap.SugaredLogger)
					bot        = ctx.Get(BotDef).(*tgbotapi.BotAPI)
					handlerMap = ctx.Get(commandDef.HandlerMapDef).(command.HandlerMap)
					handler    = ctx.Get(middleware.TelegramMiddlewareDef).(middleware2.HandlerFunc)
				)

				return source.NewTelegram(handlerMap, bot, logger, handler), nil
			},
		})
	})
}
