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
	client  *domain.Client
}

func NewRequest(
	login string,
	source domain.ClientSource,
	command string,
	args ...string,
) *Request {
	return buildRequest(login,
		source,
		command,
		args...,
	)
}

func NewClientRequest(
	client *domain.Client,
	command string,
	args ...string,
) *Request {
	if client == nil {
		return nil
	}
	req := buildRequest(client.Login, client.Source, command, args...)
	req.client = client

	return req
}

func buildRequest(login string,
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

func (r *Request) GetClient() *domain.Client {
	return r.client
}
