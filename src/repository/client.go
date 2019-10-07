package repository

import (
	"chat/src/domain"
	"github.com/go-pg/pg"
)

type ClientRepository struct {
	db *pg.DB
}

func NewClientRepository(db *pg.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (c *ClientRepository) GetByLogin(login string) (*domain.Client, error) {
	return nil, nil
}
