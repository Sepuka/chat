package command

import (
	"chat/src/command"
	"chat/src/def"
	"chat/src/def/source"
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
				return command.NewList(listPrecept), nil
			},
		})
	})
}
