package command

import (
	"chat/src/context"
	"fmt"
)

type create struct {
	precept string
}

func NewCreate(precept string) *create {
	return &create{
		precept:precept,
	}
}

func (c *create) Exec(req *context.Request) error {
	fmt.Println(`create!`)
	return nil
}

func (c *create) Precept() string {
	return c.precept
}