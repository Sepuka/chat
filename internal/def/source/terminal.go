package source

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/config"
	"github.com/sepuka/chat/internal/def"
	commandDef "github.com/sepuka/chat/internal/def/command"
	middlewareDef "github.com/sepuka/chat/internal/def/middleware"
	"github.com/sepuka/chat/internal/def/repository"
	"github.com/sepuka/chat/internal/domain"
	middleware3 "github.com/sepuka/chat/internal/middleware"
	"github.com/sepuka/chat/internal/source"
)

const (
	TerminalDef = `command.handler.terminal.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: TerminalDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					handlerMap = ctx.Get(commandDef.HandlerMapDef).(command.HandlerMap)
					clientRepo = ctx.Get(repository.ClientRepoDef).(domain.ClientRepository)
					handler    = ctx.Get(middlewareDef.TerminalMiddlewareDef).(middleware3.HandlerFunc)
				)

				return source.NewTerminal(handlerMap, clientRepo, handler), nil
			},
		})
	})
}
