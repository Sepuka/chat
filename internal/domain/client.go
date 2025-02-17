package domain

import (
	"time"

	"github.com/go-pg/pg"
)

type ClientRepository interface {
	GetByLogin(string, ClientSource) (*Client, error)
	Add(string, ClientSource) (*Client, error)
}

type Client struct {
	Id         uint64      `sql:",pk"`
	Login      string      `sql:",unique,notnull"`
	CreatedAt  time.Time   `sql:",notnull,default:now()"`
	DeletedAt  pg.NullTime `pg:",soft_delete"`
	Source     ClientSource
	Properties *ClientProperties `sql:"fk:client_id,notnull"`
}

func (c *Client) IsLimitExceeded(actual int) bool {
	if c.Properties == nil {
		return actual >= DefaultHostsLimit
	}

	return int(c.Properties.HostsLimit) <= actual
}

func (c *Client) IsTheSameUser(client *Client) bool {
	return c.Login == client.Login &&
		c.Source == client.Source
}
