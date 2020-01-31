package command

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-pg/pg"

	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
	"go.uber.org/zap"
)

const (
	cmdCreateJoomla        = `docker run -d --name %s -p %d:80 -p %d:22 --storage-opt size=2G sepuka/joomla.volatiland`
	cmdCreateEmpty         = `docker run -d --name %s -p %d:80 -p %d:22 --storage-opt size=2G sepuka/empty.volatiland`
	containerHashLength    = 12
	containerPostfixFormat = `20060102150405`
	imageJoomla            = `joomla`
	imageEmpty             = ``
)

var (
	FreePoolAreAbsent        = errors.New(`free pools are absent`)
	HostsLimitExceeded       = errors.New(`hosts limit exceeded`)
	cannotBuildContainerName = errors.New(`unknown image name`)
	availableImages          = []string{
		imageJoomla,
	}
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

func (c *Create) Exec(req *context.Request, resp *Result) error {
	var (
		pool      *domain.Pool
		client    = req.GetClient()
		trx       *pg.Tx
		hosts     []*domain.VirtualHost
		host      *domain.VirtualHost
		err       error
		remoteCmd domain.RemoteCmd
	)

	resp.Response = []byte(`internal error`)

	if client == nil {
		if client, err = c.clientRepo.Add(req.GetLogin(), req.GetSource()); err != nil {
			c.logger.Errorf(`unable to create new client %s`, req.GetFQDN())
			resp.Response = []byte(`unable to register new client`)

			return err
		}
	}

	if hosts, err = c.hostRepo.GetUsersHosts(client); err != nil {
		c.logger.Errorf(`cannot check exists client's hosts: %s`, err)
		return HostsLimitExceeded
	}
	if client.IsLimitExceeded(len(hosts)) {
		return HostsLimitExceeded
	}

	pool, host, trx, err = c.FindPool(client)
	if err != nil {
		c.logger.Errorf(`unable to find any free pool: %s`, err)
		if trx != nil {
			if rejectErr := c.rejectHost(trx); rejectErr != nil {
				c.logger.Errorf(`unable to reject trx after finding vacant pool for user %d: %s`, client.Id, err)
			}
		}
		resp.Response = []byte(`no free pool`)

		return err
	}

	c.buildPorts(pool, host)

	if remoteCmd, err = c.buildCommand(req, client, host.WebPort, host.SshPort); err != nil {
		if rejectErr := c.rejectHost(trx); rejectErr != nil {
			c.logger.Errorf(`unable to reject trx after building remote command for user %d: %s`, client.Id, err)
		}
		resp.Response = c.getAvailableImages()

		return nil
	}

	answer, err := c.cloud.Run(pool, remoteCmd)
	c.logger.Debugf(`pool #%d returned "%s" for client #%d (%s@%s)`, pool.Id, answer, client.Id, client.Login, client.Source)

	if err != nil {
		c.logger.Errorf("unable to create new virtual host: %s\nhost answered: %s", err, answer)
		if rejectErr := c.rejectHost(trx); rejectErr != nil {
			c.logger.Errorf(`unable to reject new virtual host in pool %d for user %d: %s`, pool.Id, client.Id, err)
		}

		return err
	} else {
		host.Container = string(answer[:containerHashLength])
		if err = c.hostRepo.Update(trx, host); err != nil {
			c.logger.Errorf(`error while updating virtual host %d: %s`, host.Id, err)
			if rejectErr := c.rejectHost(trx); rejectErr != nil {
				c.logger.Errorf(`unable to reject new virtual host in pool %d for user %d: %s`, pool.Id, client.Id, err)
			}
			return err
		}
		if err = c.poolRepo.Engage(pool, trx); err != nil {
			c.logger.Errorf(`cannot engage new virtual host %s`, err)
			return err
		}
	}

	resp.Response = []byte(`your new virtual host just created`)

	return nil
}

func (c *Create) FindPool(client *domain.Client) (*domain.Pool, *domain.VirtualHost, *pg.Tx, error) {
	var (
		pool *domain.Pool
		tx   *pg.Tx
		err  error
		host *domain.VirtualHost
	)

	if pool, tx, err = c.poolRepo.OccupyVacant(); err != nil {
		if err == pg.ErrNoRows {
			c.logger.Error(`unable to find any vacant pool: `, err)
			return nil, nil, nil, FreePoolAreAbsent
		}
	}

	if host, err = c.hostRepo.Add(tx, pool, client); err != nil {
		c.logger.Error(`unable to add new virtual host: `, err)
		return nil, nil, nil, err
	}

	return pool, host, tx, err
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

func (c *Create) buildCommand(req *context.Request, client *domain.Client, webPort uint16, sshPort uint16) (domain.RemoteCmd, error) {
	var (
		container = fmt.Sprintf(`%s_%d_%s`, client.Login, client.Source, time.Now().Format(containerPostfixFormat))
	)

	if len(req.GetArgs()) == 0 {
		return domain.RemoteCmd(fmt.Sprintf(cmdCreateEmpty, container, webPort, sshPort)), nil
	}

	switch req.GetArgs()[0] {
	case imageJoomla:
		return domain.RemoteCmd(fmt.Sprintf(cmdCreateJoomla, container, webPort, sshPort)), nil
	default:
		return ``, cannotBuildContainerName
	}
}

func (c *Create) buildPorts(pool *domain.Pool, host *domain.VirtualHost) {
	host.WebPort = pool.PortCnt + 1
	host.SshPort = pool.PortCnt + 2
	pool.PortCnt += 2
	pool.Workload++
}

func (c *Create) getAvailableImages() []byte {
	return bytes.NewBufferString(
		fmt.Sprintf(`available images are: %s`, strings.Join(availableImages, `,`)),
	).Bytes()
}
