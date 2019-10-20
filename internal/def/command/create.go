package command

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/def"
	"github.com/sepuka/chat/internal/def/log"
	"github.com/sepuka/chat/internal/def/repository"
	"github.com/sepuka/chat/internal/def/source"
	"github.com/sepuka/chat/internal/domain"
	"go.uber.org/zap"
)

const (
	CreateDef = `def.command.create`
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

				return command.NewCreate(pool, logger), nil
			},
		})
	})
}
