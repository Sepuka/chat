package domain

import (
	"time"
)

type ClientRepository interface {
	GetByLogin(string) (*Client, error)
}

type Client struct {
	Id        uint64    `sql:",pk"`
	Login     string    `sql:",unique,notnull"`
	CreatedAt time.Time `sql:",notnull,default:now()"`
	DeletedAt time.Time `pg:",soft_delete"`
	Source    ClientSource
}
