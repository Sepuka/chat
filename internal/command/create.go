package command

import (
	"errors"

	"github.com/go-pg/pg"

	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
	"go.uber.org/zap"
)

const (
	cmdCreate domain.RemoteCmd = `docker run -d sepuka/joomla.volatiland`
)

var (
	FreePoolAreAbsent  = errors.New(`free pools are absent`)
	HostsLimitExceeded = errors.New(`hosts limit exceeded`)
)

type Create struct {
	poolRepo   domain.PoolRepository
	hostRepo   domain.VirtualHostRepository
	clientRepo domain.ClientRepository
	logger     *zap.SugaredLogger
	cloud      domain.Cloud
}

func NewCreate(
	pool domain.PoolRepository,
	hostRepository domain.VirtualHostRepository,
	clientRepo domain.ClientRepository,
	logger *zap.SugaredLogger,
	cloud domain.Cloud,
) *Create {
	return &Create{
		poolRepo:   pool,
		hostRepo:   hostRepository,
		clientRepo: clientRepo,
		logger:     logger,
		cloud:      cloud,
	}
}

func (c *Create) Exec(req *context.Request) (*Result, error) {
	var (
		pool   *domain.Pool
		client *domain.Client
		trx    *pg.Tx
		hosts  []*domain.VirtualHost
		err    error
		result = &Result{
			Response: []byte(`internal error`),
		}
	)

	client, err = c.clientRepo.GetByLogin(req.GetLogin())
	if err != nil {
		if err == pg.ErrNoRows {
			if clientErr := c.clientRepo.Add(req.GetLogin(), req.GetSource()); clientErr != nil {
				c.logger.Errorf(`unable to create new client %s`, req.GetFQDN())
				result.Response = []byte(`unable to register new client`)

				return result, err
			}
		} else {
			c.logger.Errorf(`client %s not found: %s`, req.GetFQDN(), err)

			return result, err
		}
	}

	if hosts, err = c.hostRepo.GetUsersHosts(client); err != nil {
		c.logger.Errorf(`cannot check exists client's hosts: %s`, err)
		return result, HostsLimitExceeded
	}
	if client.IsLimitExceeded(len(hosts)) {
		return result, HostsLimitExceeded
	}

	pool, trx, err = c.FindPool(client)
	if err != nil {
		c.logger.Errorf(`unable to find any free pool: %s`, err)
		result.Response = []byte(`no free pool`)

		return result, err
	}

	answer, err := c.cloud.Run(pool, cmdCreate)
	c.logger.Debugf(`pool #%d returned "%s" for client #%d (%s@%s)`, answer, client.Id, client.Login, client.Source)

	if err != nil {
		c.logger.Errorf(`unable to create new virtual host: %s`, err)
		if rejectErr := c.rejectHost(trx); rejectErr != nil {
			c.logger.Errorf(`unable to reject new virtual host in pool %d for user %d: %s`, pool.Id, client.Id, err)
		}

		return result, err
	} else {
		if err = c.poolRepo.Engage(pool, trx); err != nil {
			c.logger.Errorf(`cannot engage new virtual host %s`, err)
			return result, err
		}
	}

	result.Response = []byte(`your new virtual host just created`)

	return result, nil
}

func (c *Create) FindPool(client *domain.Client) (*domain.Pool, *pg.Tx, error) {
	var (
		pool *domain.Pool
		trx  *pg.Tx
		err  error
	)

	if pool, trx, err = c.poolRepo.OccupyVacant(); err != nil {
		if err == pg.ErrNoRows {
			c.logger.Error(`unable to find any vacant pool: `, err)
			return nil, nil, FreePoolAreAbsent
		}
	}

	if err = c.hostRepo.Add(pool, client); err != nil {
		c.logger.Error(`unable to add new virtual host: `, err)
		return nil, nil, err
	}

	return pool, trx, err
}

func (c *Create) rejectHost(trx *pg.Tx) error {
	return trx.Rollback()
}

func (c *Create) Precept() []string {
	return []string{
		`create`,
		`/create`,
	}
}
