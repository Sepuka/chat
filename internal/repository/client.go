package repository

import (
	"time"

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
		Where(`client.login = ?`, login).
		Select()

	return client, err
}

func (c *ClientRepository) Add(login string, source domain.ClientSource) error {
	var client = &domain.Client{
		Login:     login,
		CreatedAt: time.Now(),
		DeletedAt: pg.NullTime{},
		Source:    source,
	}

	return c.db.Insert(client)
}
