package db

import (
	"chat/src/def"
	"net"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/sarulabs/di"
)

const DataBaseDef = "db.def"

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: DataBaseDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					db  *pg.DB
					cfg = ctx.Get(def.CfgDef).(def.Config)
				)

				db = pg.Connect(&pg.Options{
					User:     cfg.Database.User,
					Password: cfg.Database.Password,
					Addr:     net.JoinHostPort(cfg.Database.Host, strconv.Itoa(cfg.Database.Port)),
					Database: cfg.Database.Name,
				})

				_, err := db.Exec(`SET timezone TO 'UTC'`)

				return db, err
			},
		})
	})
}
