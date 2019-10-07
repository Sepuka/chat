package command

import (
	"chat/src/command"
	"chat/src/def"
	"chat/src/def/repository"
	"chat/src/def/source"
	"chat/src/domain"
	"github.com/sarulabs/di"
)

const (
	ListDef     = `def.command.list`
	listPrecept = `list`
)

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: ListDef,
			Tags: []di.Tag{
				{
					Name: source.CommandTagName,
					Args: nil,
				},
			},
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					clientRepo = ctx.Get(repository.ClientRepoDef).(domain.ClientRepository)
				)
				return command.NewList(listPrecept, clientRepo), nil
			},
		})
	})
}
