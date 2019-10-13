package repository

import (
	"github.com/go-pg/pg"
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/src/def"
	dbDef "github.com/sepuka/chat/src/def/db"
	"github.com/sepuka/chat/src/repository"
)

const ClientRepoDef = `repo.client.def`

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: ClientRepoDef,
			Build: func(ctx def.Context) (interface{}, error) {
				var (
					db = ctx.Get(dbDef.DataBaseDef).(*pg.DB)
				)

				return repository.NewClientRepository(db), nil
			},
		})
	})
}
