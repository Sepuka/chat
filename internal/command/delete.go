package command

import (
	"bytes"
	"fmt"
	"time"

	"github.com/sepuka/chat/internal/view"

	"github.com/go-pg/pg"
	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
	"go.uber.org/zap"
)

const (
	cmdDelete = `docker rmi %s`
)

type Delete struct {
	clientRepo domain.ClientRepository
	hostsRepo  domain.VirtualHostRepository
	poolRepo   domain.PoolRepository
	logger     *zap.SugaredLogger
	cloud      domain.Cloud
}

func NewDelete(
	clientRepo domain.ClientRepository,
	hostsRepo domain.VirtualHostRepository,
	poolRepo domain.PoolRepository,
	logger *zap.SugaredLogger,
	cloud domain.Cloud,
) *Delete {
	return &Delete{
		clientRepo: clientRepo,
		hostsRepo:  hostsRepo,
		poolRepo:   poolRepo,
		logger:     logger,
		cloud:      cloud,
	}
}

func (d *Delete) Exec(req *context.Request) (*Result, error) {
	var (
		argsNum = len(req.GetArgs())
		err     error
		client  *domain.Client
		host    *domain.VirtualHost
		tx      *pg.Tx

		result = &Result{
			Response: []byte(`internal error`),
		}
	)

	if argsNum != 1 {
		return nil, NoContainerIdError
	}

	if len(req.GetArgs()[0]) != 12 && len(req.GetArgs()[0]) != 64 {
		return nil, WrongContainerIdFormat
	}

	client, err = d.clientRepo.GetByLogin(req.GetLogin())
	if err != nil {
		if err == pg.ErrNoRows {
			result.Response = []byte(`you have not any hosts`)
			return result, nil
		} else {
			d.logger.Errorf(`client %s not found: %s`, req.GetFQDN(), err)

			return result, err
		}
	}

	host, err = d.hostsRepo.GetByContainerId(req.GetArgs()[0])
	if err != nil {
		if err == pg.ErrNoRows {
			var hosts []*domain.VirtualHost
			var availableContainers *bytes.Buffer
			if hosts, err = d.hostsRepo.GetUsersHosts(client); err != nil {
				return result, err
			}

			availableContainers = bytes.NewBufferString("Bellow available host hashes:\n" + view.NewContainersListFormatter(hosts).Format())
			result.Response = availableContainers.Bytes()
			return result, nil
		}
		return result, err
	}

	if !host.Client.IsTheSameUser(client) {
		return nil, NoHostsByContainerId
	}

	answer, err := d.cloud.Run(host.Pool, d.buildCommand(host.Container))
	d.logger.Debugf(`pool #%d returned "%s" for client #%d (%s@%s)`, host.Pool.Id, answer, client.Id, client.Login, client.Source)
	if err != nil {
		d.logger.Errorf(`unable to delete virtual host: %s (%s)`, err, answer)

		return result, err
	}

	return result, d.release(tx, host.Pool, host)
}

func (d *Delete) buildCommand(name string) domain.RemoteCmd {
	return domain.RemoteCmd(fmt.Sprintf(cmdDelete, name))
}

func (d *Delete) release(tx *pg.Tx, pool *domain.Pool, host *domain.VirtualHost) error {
	var (
		err error
	)
	tx, err = d.poolRepo.Release(pool)
	if err != nil {
		d.logger.Errorf(`unable to release the pool: %s`, err)

		return err
	}

	host.DeletedAt = pg.NullTime{Time: time.Now()}

	return d.hostsRepo.Update(tx, host)
}

func (c *Delete) Precept() []string {
	return []string{
		`delete`,
		`/delete`,
	}
}
