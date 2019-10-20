package db

import (
	"net"
	"strconv"

	"github.com/sepuka/chat/internal/config"

	"github.com/sepuka/chat/internal/def"

	"github.com/go-pg/pg"
	"github.com/sarulabs/di"
)

const DataBaseDef = "db.def"

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: DataBaseDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					db *pg.DB
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
