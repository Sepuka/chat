package cmd

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/sepuka/chat/internal/def"
	db2 "github.com/sepuka/chat/internal/def/db"
	"github.com/sepuka/chat/internal/domain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateTables)
}

var (
	generateTables = &cobra.Command{
		Use:     `db generate`,
		Example: `db generate -c /config/path`,

		RunE: func(cmd *cobra.Command, args []string) error {
			var db *pg.DB
			if err := def.Container.Fill(db2.DataBaseDef, &db); err != nil {
				return err
			}

			for _, model := range []interface{}{&domain.Pool{}, &domain.Client{}, &domain.VirtualHost{}} {
				err := db.CreateTable(model, &orm.CreateTableOptions{
					FKConstraints: true,
				})
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
)
