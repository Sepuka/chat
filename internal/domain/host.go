package domain

import (
	"github.com/go-pg/pg"
)

type VirtualHostRepository interface {
	GetUsersHosts(*Client) ([]*VirtualHost, error)
	Add(*pg.Tx, *Pool, *Client) (*VirtualHost, error)
	Update(tx *pg.Tx, host *VirtualHost) error
	GetByContainerId(string) (*VirtualHost, error)
}

type VirtualHost struct {
	Id        uint64      `sql:",pk"`
	PoolId    uint64      `sql:",notnull"`
	ClientId  uint64      `sql:",notnull"`
	CreatedAt pg.NullTime `sql:",notnull,default:now()"`
	UpdatedAt pg.NullTime `sql:",notnull,default:now()"`
	DeletedAt pg.NullTime `pg:",soft_delete"`
	Container string      `sql:",notnull"`
	Pool      *Pool       `sql:"fk:id,notnull"`
	Client    *Client     `sql:"fk:id,notnull"`
	WebPort   uint16      `sql:",notnull"`
	SshPort   uint16      `sql:",notnull"`
}
