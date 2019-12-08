package log

import (
	"errors"

	errPkg "github.com/pkg/errors"

	"go.uber.org/zap/zapcore"

	"github.com/sepuka/chat/internal/config"

	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/def"
	"go.uber.org/zap"
)

const LoggerDef = `logger.def`

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: LoggerDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					err                   error
					logger                *zap.Logger
					sugar                 *zap.SugaredLogger
					zapCfg                zap.Config
					core                  zapcore.Core
					fileEncoder           zapcore.Encoder
					telegramEncoder       zapcore.Encoder
					fileEncoderConfig     zapcore.EncoderConfig
					telegramEncoderConfig = zap.NewDevelopmentEncoderConfig()
				)

				fileSinker, closeOut, err := zap.Open(cfg.Log.Output)
				if err != nil {
					return nil, errPkg.Wrap(err, `unable to open output files`)
				}

				writeSyncer := zapcore.AddSync(fileSinker)
				telegramSyncer, err := ctx.SafeGet(LoggerSyncerDef)
				if err != nil {
					return nil, errPkg.Wrap(err, `unable to get telegram logger syncer`)
				}

				errorsOnlyLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
					return lvl == zapcore.ErrorLevel
				})
				consoleMsgLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
					if cfg.Log.Prod {
						return lvl >= zapcore.InfoLevel
					}

					return true
				})

				if cfg.Log.Prod {
					zapCfg = zap.NewProductionConfig()
					fileEncoderConfig = zap.NewProductionEncoderConfig()
				} else {
					zapCfg = zap.NewDevelopmentConfig()
					fileEncoderConfig = zap.NewDevelopmentEncoderConfig()
				}

				zapCfg.OutputPaths = []string{cfg.Log.Output}

				fileEncoder = zapcore.NewJSONEncoder(fileEncoderConfig)
				telegramEncoder = zapcore.NewConsoleEncoder(telegramEncoderConfig)
				core = zapcore.NewTee(
					zapcore.NewCore(fileEncoder, writeSyncer, consoleMsgLevel),
					zapcore.NewCore(telegramEncoder, telegramSyncer.(zapcore.WriteSyncer), errorsOnlyLevel),
				)

				logger = zap.New(core)
				sugar = logger.Sugar()
				if sugar == nil {
					closeOut()
					return nil, errors.New(`unable build sugar logger`)
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
