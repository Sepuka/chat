package repository

import (
	"time"

	"github.com/pkg/errors"

	"github.com/go-pg/pg"
	"github.com/sepuka/chat/internal/domain"
)

type ClientRepository struct {
	db *pg.DB
}

func NewClientRepository(db *pg.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (c *ClientRepository) GetByLogin(login string) (*domain.Client, error) {
	var (
		client = &domain.Client{}
		err    error
	)
	err = c.
		db.
		Model(client).
		Relation(`Properties`).
		Where(`client.login = ?`, login).
		Select()

	return client, err
}

func (c *ClientRepository) Add(login string, source domain.ClientSource) error {
	var (
		err    error
		client = &domain.Client{
			Login:     login,
			CreatedAt: time.Now(),
			DeletedAt: pg.NullTime{},
			Source:    source,
		}
		properties = &domain.ClientProperties{
			ClientId:   0,
			HostsLimit: domain.DefaultHostsLimit,
		}
	)

	trx, err := c.db.Begin()
	if err != nil {
		return err
	}

	err = c.db.Insert(client)
	if err != nil {
		if rollbackErr := trx.Rollback(); rollbackErr != nil {
			return errors.Wrapf(rollbackErr, `cannot rollback transaction after client inserting, error '%s'`, err.Error())
		}
		return err
	}

	properties.ClientId = client.Id
	err = c.db.Insert(properties)
	if err != nil {
		if rollbackErr := trx.Rollback(); rollbackErr != nil {
			return errors.Wrapf(rollbackErr, `cannot rollback transaction after properties inserting, error '%s'`, err.Error())
		}
		return err
	}

	return trx.Commit()
}
