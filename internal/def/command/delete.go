package command

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/config"
	"github.com/sepuka/chat/internal/def"
	"github.com/sepuka/chat/internal/def/cloud"
	"github.com/sepuka/chat/internal/def/log"
	"github.com/sepuka/chat/internal/def/repository"
	"github.com/sepuka/chat/internal/domain"
	"go.uber.org/zap"
)

const (
	DeleteDef = `def.command.delete`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: DeleteDef,
			Tags: []di.Tag{
				{
					Name: ExecutorDef,
					Args: nil,
				},
			},
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					poolRepo   = def.Container.Get(repository.PoolRepoDef).(domain.PoolRepository)
					hostRepo   = def.Container.Get(repository.HostRepoDef).(domain.VirtualHostRepository)
					clientRepo = def.Container.Get(repository.ClientRepoDef).(domain.ClientRepository)
					logger     = def.Container.Get(log.LoggerDef).(*zap.SugaredLogger)
					sshClient  = def.Container.Get(cloud.SshCloudDef).(domain.Cloud)
				)

				return command.NewDelete(clientRepo, hostRepo, poolRepo, logger, sshClient), nil
			},
		})
	})
}
