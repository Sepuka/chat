package source

import (
	"chat/src/command"
	"errors"
)

var (
	emptyInstruction = errors.New(`got empty instruction`)
	unknownInstruction = errors.New(`got unknown instruction`)
)

type Terminal struct {
	commands map[string]command.Executor
}

func NewTerminal(
	commandsMap map[string]command.Executor,
) *Terminal {
	return &Terminal{
		commands: commandsMap,
	}
}

func (hosting *Terminal) Execute(args []string) error {
	if len(args) == 0 {
		return emptyInstruction
	}
	instr := args[0]
	if f, ok := hosting.commands[instr]; ok {
		return f.Exec()
	}

	return unknownInstruction
}
