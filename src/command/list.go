package command

import "fmt"

type list struct {
	precept string
}

func NewList(precept string) *list {
	return &list{
		precept:precept,
	}
}

func (l *list) Exec() {
	fmt.Println(`list!`)
}

func (c *list) Precept() string {
	return c.precept
}