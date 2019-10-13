package repository

import (
	"github.com/go-pg/pg"
	"github.com/sepuka/chat/src/domain"
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
		Where(`client.login = ?`, login).Select()

	return client, err
}
