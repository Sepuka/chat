package command

import "github.com/sepuka/chat/internal/context"

type Result struct {
	Msg string
}

type Executor interface {
	Exec(*context.Request) (*Result, error)
}

type Preceptable interface {
	Precept() string
}
