package repository

import (
	"github.com/go-pg/pg"
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/internal/def"
	dbDef "github.com/sepuka/chat/internal/def/db"
	"github.com/sepuka/chat/internal/repository"
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
