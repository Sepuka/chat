package command

import (
	"errors"

	"github.com/sepuka/chat/internal/view"

	"go.uber.org/zap"

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

func (l *List) Exec(req *context.Request, resp *Result) error {
	var (
		formatter = view.NewShortHostsListFormatter(req.GetSource())
		client    = req.GetClient()
	)

	resp.Response = []byte(`you're have not any hosts`)

	if client == nil {
		return nil
	}

	hosts, err := l.hostsRepo.GetUsersHosts(client)
	if err != nil {
		l.logger.Error(
			`db error`,
			zap.Error(err),
			zap.String(`user`, req.GetLogin()),
			zap.String(`command`, req.GetCommand()),
			zap.Strings(`args`, req.GetArgs()),
		)
		return errors.New(`some error occurred`)
	}

	resp.Response = []byte(formatter.Format(hosts))

	return nil
}

func (l *List) Precept() []string {
	return []string{
		`list`,
		`/list`,
	}
}
