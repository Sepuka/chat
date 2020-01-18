package middleware

import (
	"time"

	"github.com/sepuka/chat/internal/def"
	log2 "github.com/sepuka/chat/internal/def/log"
	"go.uber.org/zap"

	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/context"
)

func Duration(next HandlerFunc) HandlerFunc {
	return func(exec command.Executor, req *context.Request, res *command.Result, err error) {
		var start = time.Now()
		next(exec, req, res, err)
		var duration = time.Since(start)
		var log = def.Container.Get(log2.LoggerDef).(*zap.SugaredLogger)
		log.Infof(`duration %s is %s`, req.GetCommand(), duration.String())
	}
}
