package command

import (
	"chat/src/context"
	"fmt"
)

type list struct {
	precept string
}

func NewList(precept string) *list {
	return &list{
		precept:precept,
	}
}

func (l *list) Exec(req *context.Request) error {
	fmt.Println(`list!`)
	return nil
}

func (c *list) Precept() string {
	return c.precept
}