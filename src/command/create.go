package command

import "fmt"

type create struct {
	precept string
}

func NewCreate(precept string) *create {
	return &create{
		precept:precept,
	}
}

func (c *create) Exec() {
	fmt.Println(`create!`)
}

func (c *create) Precept() string {
	return c.precept
}