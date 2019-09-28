package command

import (
	"chat/src/command"
	"chat/src/def"
	"chat/src/def/source"
	"github.com/sarulabs/di"
)

const (
	CreateDef     = `def.command.create`
	createPrecept = `create`
)

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: CreateDef,
			Tags: []di.Tag{
				{
					Name: source.CommandTagName,
					Args: nil,
				},
			},
			Build: func(ctx def.Context) (interface{}, error) {
				return command.NewCreate(createPrecept), nil
			},
		})
	})
}
