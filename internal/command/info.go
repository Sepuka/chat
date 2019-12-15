package command

import (
	"errors"

	"github.com/sepuka/chat/internal/format"

	"github.com/go-pg/pg"

	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
	"go.uber.org/zap"
)

var (
	NoContainerIdError     = errors.New(`there is not any container id`)
	WrongContainerIdFormat = errors.New(`expected 12 or 64 symbols of container id`)
	NoHostsByContainerId   = errors.New(`hosts with specified container id does not exists`)
)

type Info struct {
	clientRepo domain.ClientRepository
	hostsRepo  domain.VirtualHostRepository
	logger     *zap.SugaredLogger
}

func NewInfo(
	clientRepo domain.ClientRepository,
	hostsRepo domain.VirtualHostRepository,
	logger *zap.SugaredLogger,
) *Info {
	return &Info{
		clientRepo: clientRepo,
		hostsRepo:  hostsRepo,
		logger:     logger,
	}
}

func (l *Info) Exec(req *context.Request) (*Result, error) {
	var (
		argsNum = len(req.GetArgs())
		err     error
		client  *domain.Client
		host    *domain.VirtualHost
		result  = &Result{
			Response: []byte(`internal error`),
		}
	)

	if argsNum != 1 {
		return nil, NoContainerIdError
	}

	if len(req.GetArgs()[0]) != 12 && len(req.GetArgs()[0]) != 64 {
		return nil, WrongContainerIdFormat
	}

	client, err = l.clientRepo.GetByLogin(req.GetLogin())
	if err != nil {
		if err == pg.ErrNoRows {
			result.Response = []byte(`you have not any hosts`)
			return result, nil
		}
		return result, err
	}

	host, err = l.hostsRepo.GetByContainerId(req.GetArgs()[0])
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, NoHostsByContainerId
		}
		return result, err
	}

	if !host.Client.IsTheSameUser(client) {
		return nil, NoHostsByContainerId
	}

	formatter := format.NewInfoFormatter(host)
	result.Response = formatter.Format()

	return result, nil
}

func (l *Info) Precept() []string {
	return []string{
		`info`,
		`/info`,
	}
}
