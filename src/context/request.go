package context

type Request struct {
	login   string
	command string
	args    []string
}

func NewRequest(
	login string,
	command string,
	args ...string,
) *Request {
	return &Request{
		login:   login,
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
