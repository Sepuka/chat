package context

import (
	"fmt"

	"github.com/sepuka/chat/internal/domain"
)

const (
	maxCmdLength = 32
)

type Request struct {
	login   string
	source  domain.ClientSource
	command string
	args    []string
}

func NewRequest(
	login string,
	source domain.ClientSource,
	command string,
	args ...string,
) *Request {
	if len(command) >= maxCmdLength {
		command = command[:maxCmdLength]
	}
	return &Request{
		login:   login,
		source:  source,
		command: command,
		args:    args,
	}
}

func (r *Request) GetLogin() string {
	return r.login
}

func (r *Request) GetCommand() string {
	return r.command
}

func (r *Request) GetArgs() []string {
	return r.args
}

func (r *Request) GetSource() domain.ClientSource {
	return r.source
}

func (r *Request) GetFQDN() string {
	return fmt.Sprintf(`%s@%d`, r.login, r.source)
}
