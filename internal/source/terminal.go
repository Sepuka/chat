package source

import (
	"errors"

	"github.com/sepuka/chat/internal/middleware"

	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
)

var (
	unknownInstruction = errors.New(`got unknown instruction`)
)

type Terminal struct {
	commands   command.HandlerMap
	clientRepo domain.ClientRepository
	handler    middleware.HandlerFunc
}

func NewTerminal(
	commandsMap command.HandlerMap,
	clientRepo domain.ClientRepository,
	handler middleware.HandlerFunc,
) *Terminal {
	return &Terminal{
		commands:   commandsMap,
		clientRepo: clientRepo,
		handler:    handler,
	}
}

func (src *Terminal) Execute(req *context.Request) (*command.Result, error) {
	if finalHandler, ok := src.commands[req.GetCommand()]; ok {
		var (
			resp = &command.Result{}
			err  error
		)

		err = src.handler(finalHandler, req, resp)

		return resp, err
	}

	return nil, unknownInstruction
}
