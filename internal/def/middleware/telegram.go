package middleware

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/config"
	"github.com/sepuka/chat/internal/def"
	"github.com/sepuka/chat/internal/middleware"
)

const (
	TelegramMiddlewareDef = `middleware.telegram.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: TelegramMiddlewareDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					terminalMiddleware = []func(handlerFunc middleware.HandlerFunc) middleware.HandlerFunc{
						middleware.Panic,
						middleware.Duration,
					}
				)

				return middleware.BuildHandlerChain(terminalMiddleware), nil
			},
		})
	})
}
