package middleware

import (
	"log"
	"time"

	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/context"
)

func Duration(next HandlerFunc) HandlerFunc {
	return func(exec command.Executor, req *context.Request, res *command.Result, err error) {
		var start = time.Now()
		next(exec, req, res, err)
		var duration = time.Since(start)
		log.Println(`duration %s`, duration.String())
	}
}
