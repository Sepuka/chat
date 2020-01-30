package command

import (
	"errors"

	"github.com/sepuka/chat/internal/view"

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
	hostsRepo domain.VirtualHostRepository
	logger    *zap.SugaredLogger
}

func NewInfo(
	hostsRepo domain.VirtualHostRepository,
	logger *zap.SugaredLogger,
) *Info {
	return &Info{
		hostsRepo: hostsRepo,
		logger:    logger,
	}
}

func (l *Info) Exec(req *context.Request, resp *Result) error {
	var (
		argsNum = len(req.GetArgs())
		err     error
		client  *domain.Client
		host    *domain.VirtualHost
	)

	resp.Response = []byte(`internal error`)
	if argsNum != 1 {
		return NoContainerIdError
	}

	if len(req.GetArgs()[0]) != 12 && len(req.GetArgs()[0]) != 64 {
		return WrongContainerIdFormat
	}

	client = req.GetClient()
	if client == nil {
		resp.Response = []byte(`you have not any hosts`)
		return nil
	}

	host, err = l.hostsRepo.GetByContainerId(req.GetArgs()[0])
	if err != nil {
		if err == pg.ErrNoRows {
			return NoHostsByContainerId
		}
		return err
	}

	if !host.Client.IsTheSameUser(client) {
		return NoHostsByContainerId
	}

	formatter := view.NewInfoFormatter(host)
	resp.Response = formatter.Format()

	return nil
}

func (l *Info) Precept() []string {
	return []string{
		`info`,
		`/info`,
	}
}
