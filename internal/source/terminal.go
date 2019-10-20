package source

import (
	"errors"

	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
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

func (src *Terminal) Execute(req *context.Request) (*command.Result, error) {
	if f, ok := src.commands[req.GetCommand()]; ok {
		return f.Exec(req)
	}

	return nil, unknownInstruction
}
