package middleware

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/context"
)

func Panic(next HandlerFunc) HandlerFunc {
	return func(exec command.Executor, req *context.Request, res *command.Result, err error) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("panic: %s\n"+
					"command `%s` user `%s`\n"+
					"stacktrace from panic: %s\n",
					r, req.GetCommand(), req.GetLogin(), string(debug.Stack()))
				err = errors.New(`internal error`)
			}
		}()

		next(exec, req, res, err)
	}
}
