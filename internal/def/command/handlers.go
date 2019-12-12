package command

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/config"
	"github.com/sepuka/chat/internal/def"
)

const (
	HandlerMapDef = `handler.map.def`
	ExecutorDef   = `hosting.command.tag`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: HandlerMapDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					handlers   = def.GetByTag(ExecutorDef)
					handlerMap = make(command.HandlerMap, len(handlers))
					precept    string
				)

				for _, cmd := range handlers {
					for _, precept = range cmd.(command.Preceptable).Precept() {
						handlerMap[precept] = cmd.(command.Executor)
					}
				}

				return handlerMap, nil
			},
		})
	})
}
