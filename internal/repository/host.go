package repository

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/sepuka/chat/internal/domain"
)

type VirtualHostRepository struct {
	db *pg.DB
}

func NewVirtualHostRepository(db *pg.DB) *VirtualHostRepository {
	return &VirtualHostRepository{db: db}
}

func (r *VirtualHostRepository) GetUsersHosts(client *domain.Client) ([]*domain.VirtualHost, error) {
	var (
		hosts []*domain.VirtualHost
		err   error
	)
	err = r.
		db.
		Model(&hosts).
		Where(`client_id = ?`, client.Id).Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return []*domain.VirtualHost{}, nil
		}
	}

	return nil, err
}

func (r *VirtualHostRepository) Add(pool *domain.Pool, client *domain.Client) error {
	var host = &domain.VirtualHost{
		PoolId:    pool.Id,
		ClientId:  client.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: pg.NullTime{},
	}

	return r.db.Insert(host)
}
