package command

import (
	"chat/src/context"
	"chat/src/domain"
	"fmt"
	"github.com/go-pg/pg"
)

type List struct {
	precept string
	clientRepo domain.ClientRepository
}

func NewList(
	precept string,
	clientRepo domain.ClientRepository,
	) *List {
	return &List{
		precept:precept,
		clientRepo:clientRepo,
	}
}

func (l *List) Exec(req *context.Request) error {
	client, err := l.getClient(req.Login)
	if err != nil {
		return err
	}

	if client != nil {
		fmt.Printf(`Client "%s" wants to run command LIST`, client.Login)
	} else {
		fmt.Printf(`Unknown client wants to run command LIST`)
	}
	return nil
}

func (l *List) Precept() string {
	return l.precept
}

func (l *List) getClient(login string) (*domain.Client, error) {
	client, err := l.clientRepo.GetByLogin(login)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil ,err
	}

	return client, nil
}
