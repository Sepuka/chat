package domain

import "errors"

const (
	createCmd = `create`
	listCmd   = `list`
)

type Command struct {
	Value string
}

func (c *Command) String() string {
	return c.Value
}

func NewCommand(cmd string) (*Command, error) {
	switch cmd {
	case createCmd:
		return &Command{
			Value: cmd,
		}, nil
	case listCmd:
		return &Command{
			Value: cmd,
		}, nil
	default:
		return nil, errors.New(`Invalid incoming command ` + cmd)
	}
}
