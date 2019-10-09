package repository

import (
	"chat/src/domain"
	"github.com/go-pg/pg"
)

type PoolRepository struct {
	db *pg.DB
}

func NewPoolRepository(db *pg.DB) *PoolRepository {
	return &PoolRepository{db: db}
}

func (r *PoolRepository) GetFreePool() (*domain.Pool, error) {
	var (
		pool *domain.Pool
		err  error
	)
	err = r.
		db.
		Model(pool).
		Where(`pool.active = true`).
		Order(`workload desc`).
		Limit(1).
		Select()

	return pool, err
}
