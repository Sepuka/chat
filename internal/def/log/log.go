package log

import (
	"errors"

	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/def"
	"go.uber.org/zap"
)

const LoggerDef = `logger.def`

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: LoggerDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					logger *zap.Logger
					sugar  *zap.SugaredLogger
					err    error
				)

				if cfg.Log.Prod {
					logger, err = zap.NewProduction()
				}
				if logger, err = zap.NewDevelopment(); err != nil {
					return nil, err
				}

				sugar = logger.Sugar()
				if sugar == nil {
					return sugar, errors.New(`unable build sugar logger`)
				}

				return sugar, err
			},
			Close: func(obj interface{}) error {
				logger := obj.(*zap.SugaredLogger)
				return logger.Sync()
			},
		})
	})
}
