package command

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/config"
	"github.com/sepuka/chat/internal/def"
	"github.com/sepuka/chat/internal/def/log"
	"github.com/sepuka/chat/internal/def/repository"
	"github.com/sepuka/chat/internal/def/source"
	"github.com/sepuka/chat/internal/domain"
	"go.uber.org/zap"
)

const (
	ListDef = `def.command.list`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
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
					hostsRepo  = ctx.Get(repository.HostRepoDef).(domain.VirtualHostRepository)
					logger     = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
				)
				return command.NewList(clientRepo, hostsRepo, logger), nil
			},
		})
	})
}
