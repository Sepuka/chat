package command

import (
	"github.com/sepuka/chat/src/command"
	"github.com/sepuka/chat/src/def"
	"github.com/sepuka/chat/src/def/repository"
	"github.com/sepuka/chat/src/def/source"
	"github.com/sepuka/chat/src/domain"
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
