package command

import (
	"github.com/sepuka/chat/src/command"
	"github.com/sepuka/chat/src/def"
	"github.com/sepuka/chat/src/def/log"
	"github.com/sepuka/chat/src/def/repository"
	"github.com/sepuka/chat/src/def/source"
	"github.com/sepuka/chat/src/domain"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
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
				var (
					pool   = def.Container.Get(repository.PoolRepoDef).(domain.PoolRepository)
					logger = def.Container.Get(log.LoggerDef).(*zap.SugaredLogger)
				)

				return command.NewCreate(createPrecept, pool, logger), nil
			},
		})
	})
}
