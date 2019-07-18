package domain

import (
	"net"
	"time"
)

type Pool struct {
	Id        uint64    `sql:",pk"`
	Address   net.IP    `sql:",notnull"`
	CreatedAt time.Time `sql:",notnull,default:now()"`
	UpdatedAt time.Time `sql:",notnull,default:now()"`
	DeletedAt time.Time `pg:",soft_delete"`
	Active    bool      `sql:",notnull,default:0"`
	Workload  uint64    `sql:",notnull,default:0"`
}
