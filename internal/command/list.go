package command

import (
	"errors"

	"github.com/sepuka/chat/internal/format"

	"go.uber.org/zap"

	"github.com/go-pg/pg"
	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
)

type List struct {
	clientRepo domain.ClientRepository
	hostsRepo  domain.VirtualHostRepository
	logger     *zap.SugaredLogger
}

func NewList(
	clientRepo domain.ClientRepository,
	hostsRepo domain.VirtualHostRepository,
	logger *zap.SugaredLogger,
) *List {
	return &List{
		clientRepo: clientRepo,
		hostsRepo:  hostsRepo,
		logger:     logger,
	}
}

func (l *List) Exec(req *context.Request) (*Result, error) {
	var (
		formatter = format.NewShortHostsListFormatter(req.GetSource())
	)
	client, err := l.getClient(req.GetLogin())
	if err != nil {
		return nil, err
	}
	if client != nil {
		hosts, err := l.hostsRepo.GetUsersHosts(client)
		if err != nil {
			l.logger.Error(
				`db error`,
				zap.Error(err),
				zap.String(`user`, req.GetLogin()),
				zap.String(`command`, req.GetCommand()),
				zap.Strings(`args`, req.GetArgs()),
			)
			return nil, errors.New(`some error occurred`)
		}

		return &Result{
			Response: []byte(formatter.Format(hosts)),
		}, nil
	}

	return &Result{
		Response: []byte(`you're have not any hosts`),
	}, nil
}

func (l *List) Precept() string {
	return `list`
}

func (l *List) getClient(login string) (*domain.Client, error) {
	client, err := l.clientRepo.GetByLogin(login)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return client, nil
}
