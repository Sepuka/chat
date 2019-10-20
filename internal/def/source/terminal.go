package source

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/def"
	"github.com/sepuka/chat/internal/def/repository"
	"github.com/sepuka/chat/internal/domain"
	"github.com/sepuka/chat/internal/source"
)

const (
	TerminalDef = `command.handler.terminal.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: TerminalDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					handlers   = def.GetByTag(CommandTagName)
					handlerMap = make(map[string]command.Executor, len(handlers))
					clientRepo = ctx.Get(repository.ClientRepoDef).(domain.ClientRepository)
				)

				for _, cmd := range handlers {
					precept := cmd.(command.Preceptable)
					handlerMap[precept.Precept()] = cmd.(command.Executor)
				}

				return source.NewTerminal(handlerMap, clientRepo), nil
			},
		})
	})
}
