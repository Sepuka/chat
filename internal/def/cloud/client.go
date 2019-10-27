package cloud

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/cloud"
	"github.com/sepuka/chat/internal/config"
	"github.com/sepuka/chat/internal/def"
)

const (
	SshCloudDef = `cloud.ssh.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: SshCloudDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					clientBuilder = cloud.NewClientBuilder(cfg)
				)

				return cloud.NewCloud(clientBuilder, cfg), nil
			},
		})
	})
}
