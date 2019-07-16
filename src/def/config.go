package def

import (
	"github.com/sarulabs/di"
	"github.com/stevenroose/gonfig"
)

const CfgDef = "config"

type (
	httpClient struct {
		Proxy string `id:"proxy"`
	}

	telegram struct {
		Token string `id:"token"`
	}

	Config struct {
		Path string `id:"path"`
		HttpClient httpClient `id:"http"`
		Telegram telegram  `id:"telegram"`
	}
)

func init() {
	Register(func(builder *di.Builder, cfg Config) error {
		return builder.Add(di.Def{
			Name: CfgDef,
			Build: func(ctx Context) (interface{}, error) {
				err := gonfig.Load(&cfg, gonfig.Conf{
					FileDefaultFilename: cfg.Path,
					FlagIgnoreUnknown:   true,
					FlagDisable:         true,
					EnvDisable:          true,
				})
				if err != nil {
					return nil, err
				}

				return cfg, nil
			},
		})
	})
}
