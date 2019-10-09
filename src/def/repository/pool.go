package repository

import (
	"chat/src/def"
	dbDef "chat/src/def/db"
	"chat/src/repository"
	"github.com/go-pg/pg"
	"github.com/sarulabs/di"
)

const PoolRepoDef = `repo.pool.def`

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: PoolRepoDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					db = ctx.Get(dbDef.DataBaseDef).(*pg.DB)
				)

				return repository.NewPoolRepository(db), nil
			},
		})
	})
}
