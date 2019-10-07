package command

import "chat/src/context"

type Executor interface {
	Exec(*context.Request) error
}

type Preceptable interface {
	Precept() string
}
