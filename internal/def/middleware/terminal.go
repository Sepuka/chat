package middleware

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/config"
	"github.com/sepuka/chat/internal/def"
	"github.com/sepuka/chat/internal/def/repository"
	"github.com/sepuka/chat/internal/domain"
	"github.com/sepuka/chat/internal/middleware"
)

const (
	TerminalMiddlewareDef = `middleware.terminal.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: TerminalMiddlewareDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					clientRepo         = ctx.Get(repository.ClientRepoDef).(domain.ClientRepository)
					clientMiddleware   = middleware.NewClientMiddleware(clientRepo)
					terminalMiddleware = []func(handlerFunc middleware.HandlerFunc) middleware.HandlerFunc{
						middleware.Panic,
						middleware.Duration,
						clientMiddleware.ClientHandler,
					}
				)

				return middleware.BuildHandlerChain(terminalMiddleware), nil
			},
		})
	})
}
