package command

import "github.com/sepuka/chat/src/context"

type Executor interface {
	Exec(*context.Request) error
}

type Preceptable interface {
	Precept() string
}
