package def

import (
	"github.com/sarulabs/di"
	"github.com/stevenroose/gonfig"
)

const CfgDef = "config"

type (
	httpClient struct {
		Proxy   string
	}

	telegram struct {
		Token string
		Debug bool `default:false`
	}

	database struct {
		User     string
		Password string
		Name     string
		Host     string
		Port     int
	}

	log struct {
		Prod bool
	}

	Config struct {
		Path       string
		HttpClient httpClient `id:"http"`
		Telegram   telegram
		Database   database
		Log        log
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
