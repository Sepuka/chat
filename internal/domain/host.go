package domain

import (
	"time"

	"github.com/go-pg/pg"
)

type VirtualHostRepository interface {
	GetUsersHosts(*Client) ([]*VirtualHost, error)
	Add(*Pool, *Client) error
}

type VirtualHost struct {
	Id        uint64      `sql:",pk"`
	PoolId    uint64      `sql:",notnull"`
	ClientId  uint64      `sql:",notnull"`
	CreatedAt time.Time   `sql:",notnull,default:now()"`
	UpdatedAt time.Time   `sql:",notnull,default:now()"`
	DeletedAt pg.NullTime `pg:",soft_delete"`
	Pool      *Pool       `sql:"fk:id,notnull"`
	Client    *Client     `sql:"fk:id,notnull"`
}
