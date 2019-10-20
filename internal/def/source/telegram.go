package source

import (
	"errors"
	"fmt"
	"net/http"

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
	TelegramDef    = `command.handler.telegram.def`
	BotDef         = `command.source.telegram.def`
	CommandTagName = `hosting.command.tag`
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
