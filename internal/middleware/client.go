package middleware

import (
	"github.com/go-pg/pg"
	"github.com/sepuka/chat/internal/command"
	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
)

type Client struct {
	clientRepo domain.ClientRepository
}

func NewClientMiddleware(clientRepo domain.ClientRepository) *Client {
	return &Client{
		clientRepo: clientRepo,
	}
}

func (h *Client) ClientHandler(next HandlerFunc) HandlerFunc {
	return func(exec command.Executor, req *context.Request, res *command.Result) error {
		client, err := h.clientRepo.GetByLogin(req.GetLogin())
		if err != nil {
			if err != pg.ErrNoRows {
				return err
			}
			return next(exec, req, res)
		}

		return next(exec, context.NewClientRequest(client, req.GetCommand(), req.GetArgs()...), res)
	}
}
