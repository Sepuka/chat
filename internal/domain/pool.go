package domain

import (
	"net"

	"github.com/go-pg/pg"
)

type PoolRepository interface {
	OccupyVacant() (*Pool, *pg.Tx, error)
	Engage(*Pool, *pg.Tx) error
	Release(*Pool) (*pg.Tx, error)
}

type Pool struct {
	Id        uint64      `sql:",pk"`
	Address   net.IP      `sql:",notnull"`
	CreatedAt pg.NullTime `sql:",notnull,default:now()"`
	UpdatedAt pg.NullTime `sql:",notnull,default:now()"`
	DeletedAt pg.NullTime `pg:",soft_delete"`
	Active    bool        `sql:",notnull,default:false"`
	Workload  uint64      `sql:",notnull,default:0"`
	Secret    string      `sql:",notnull"`
}
