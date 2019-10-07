package context

type Request struct {
	Login   string
	Command string
	Args    []string
}

func NewRequest(
	login string,
	command string,
	args ...string,
) *Request {
	return &Request{
		Login:   login,
		Command: command,
		Args:    args,
	}
}
