package telegram

import (
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/src/def"
	httpClient "github.com/sepuka/chat/src/def/http"
)

const BotDef = "telegram"

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: BotDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					client = ctx.Get(httpClient.ClientDef).(*http.Client)
					cfg = ctx.Get(def.CfgDef).(def.Config)
				)

				bot, err := tgbotapi.NewBotAPIWithClient(cfg.Telegram.Token, client)
				if err != nil {
					return nil, err
				}

				bot.Debug = cfg.Telegram.Debug

				return bot, nil
			},
		})
	})
}

