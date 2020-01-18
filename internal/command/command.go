package command

import (
	"github.com/sepuka/chat/internal/context"
)

type HandlerMap map[string]Executor

type Result struct {
	Code     int
	Response []byte
}

type Executor interface {
	Exec(*context.Request, *Result) error
}

type Preceptable interface {
	Precept() []string
}
