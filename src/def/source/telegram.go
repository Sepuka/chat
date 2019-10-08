package source

import (
	"chat/src/command"
	"chat/src/def"
	httpClient "chat/src/def/http"
	"chat/src/def/log"
	"chat/src/source"
	"errors"
	"fmt"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	TelegramDef    = `command.handler.telegram.def`
	BotDef         = `command.source.telegram.def`
	CommandTagName = `hosting.command.tag`
)

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: BotDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					client = ctx.Get(httpClient.ClientDef).(*http.Client)
					cfg    = ctx.Get(def.CfgDef).(def.Config)
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

	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: TelegramDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					handlers   = def.GetByTag(CommandTagName)
					handlerMap = make(map[string]command.Executor, len(handlers))
					bot        = ctx.Get(BotDef).(*tgbotapi.BotAPI)
					logger     = def.Container.Get(log.LoggerDef).(*zap.SugaredLogger)
				)

				for _, cmd := range handlers {
					precept := cmd.(command.Preceptable)
					handlerMap[precept.Precept()] = cmd.(command.Executor)
				}

				return source.NewTelegram(handlerMap, bot, logger), nil
			},
		})
	})
}
