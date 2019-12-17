package repository

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	"github.com/sepuka/chat/internal/domain"
)

type PoolRepository struct {
	db *pg.DB
}

func NewPoolRepository(db *pg.DB) *PoolRepository {
	return &PoolRepository{db: db}
}

func (r *PoolRepository) OccupyVacant() (*domain.Pool, *pg.Tx, error) {
	var (
		pool = &domain.Pool{}
		trx  *pg.Tx
		err  error
	)

	trx, err = r.db.Begin()
	if err != nil {
		return nil, nil, err
	}

	err = r.
		db.
		Model(pool).
		Where(`pool.active = ?`, true).
		Order(`workload DESC`).
		Limit(1).
		For(`UPDATE SKIP LOCKED`).
		Select()

	return pool, trx, err
}

func (r *PoolRepository) Engage(pool *domain.Pool, trx *pg.Tx) error {
	pool.Workload++
	pool.UpdatedAt = time.Now()
	updateError := r.db.Update(pool)
	if updateError != nil {
		if trxError := trx.Rollback(); trxError != nil {
			return errors.Wrapf(trxError, `rollback transaction error (update error %s)`, updateError)
		}
		return errors.Wrap(updateError, `pool update error`)
	}

	return trx.Commit()
}

func (r *PoolRepository) Release(pool *domain.Pool) (*pg.Tx, error) {
	var (
		tx  *pg.Tx
		err error
	)

	tx, err = r.db.Begin()
	if err != nil {
		return nil, err
	}

	pool.Workload--
	pool.UpdatedAt = time.Now()

	return tx, tx.Update(pool)
}
