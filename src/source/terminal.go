package source

import (
	"chat/src/command"
	"chat/src/context"
	"chat/src/domain"
	"errors"
)

var (
	unknownInstruction = errors.New(`got unknown instruction`)
)

type Terminal struct {
	commands   map[string]command.Executor
	clientRepo domain.ClientRepository
}

func NewTerminal(
	commandsMap map[string]command.Executor,
	clientRepo domain.ClientRepository,
) *Terminal {
	return &Terminal{
		commands:   commandsMap,
		clientRepo: clientRepo,
	}
}

func (src *Terminal) Execute(req *context.Request) error {
	if f, ok := src.commands[req.GetCommand()]; ok {
		return f.Exec(req)
	}

	return unknownInstruction
}
