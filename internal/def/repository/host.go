package repository

import (
	"github.com/go-pg/pg"
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/def"
	dbDef "github.com/sepuka/chat/internal/def/db"
	"github.com/sepuka/chat/internal/repository"
)

const HostRepoDef = `repo.host.def`

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: HostRepoDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					db = ctx.Get(dbDef.DataBaseDef).(*pg.DB)
				)

				return repository.NewVirtualHostRepository(db), nil
			},
		})
	})
}
