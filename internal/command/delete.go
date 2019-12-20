package command

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/sepuka/chat/internal/view"

	"github.com/go-pg/pg"
	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
	"go.uber.org/zap"
)

const (
	cmdDelete = `docker rm -fv %s`
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
		d.logger.Errorf("unable to delete virtual host at pool #%d: %s\n%s", host.Pool.Id, err, answer)

		return result, err
	}

	if tx, err = d.release(host.Pool, host); err != nil {
		d.logger.Errorf(`error while rollback release host #%d at pool #%d: %s`, host.Id, host.Pool.Id, err)
		if err = tx.Rollback(); err != nil {
			d.logger.Errorf(`error while rollback delete host #%d at pool #%d: %s`, host.Id, host.Pool.Id, err)
		}
		return result, errors.New(`cannot delete host`)
	}

	if err = tx.Commit(); err != nil {
		d.logger.Errorf(`error while commit deleting host #%d at pool #%d: %s`, host.Id, host.Pool.Id, err)
		return result, errors.New(`cannot delete host`)
	}

	result.Response = []byte(`host was deleted`)

	return result, nil
}

func (d *Delete) buildCommand(name string) domain.RemoteCmd {
	return domain.RemoteCmd(fmt.Sprintf(cmdDelete, strings.TrimSpace(name)))
}

func (d *Delete) release(pool *domain.Pool, host *domain.VirtualHost) (*pg.Tx, error) {
	var (
		err error
		tx  *pg.Tx
	)
	tx, err = d.poolRepo.Release(pool)
	if err != nil {
		d.logger.Errorf(`unable to release the pool: %s`, err)

		return nil, err
	}

	return tx, tx.Delete(host)
}

func (d *Delete) Precept() []string {
	return []string{
		`delete`,
		`/delete`,
	}
}
