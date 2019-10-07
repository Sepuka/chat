package source

import (
	"chat/src/command"
	"chat/src/def"
	"chat/src/def/repository"
	"chat/src/domain"
	"chat/src/source"
	"github.com/sarulabs/di"
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
