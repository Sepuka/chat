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
		Column(`virtual_host.*`).
		Relation(`Pool`).
		Relation(`Client`).
		Where(`client_id = ?`, client.Id).Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return []*domain.VirtualHost{}, nil
		}
	}

	return hosts, err
}

func (r *VirtualHostRepository) Add(tx *pg.Tx, pool *domain.Pool, client *domain.Client) (*domain.VirtualHost, error) {
	var host = &domain.VirtualHost{
		PoolId:    pool.Id,
		ClientId:  client.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: pg.NullTime{},
	}

	return host, tx.Insert(host)
}

func (r *VirtualHostRepository) Update(tx *pg.Tx, host *domain.VirtualHost) error {
	host.UpdatedAt = time.Now()

	return tx.Update(host)
}

func (r *VirtualHostRepository) GetByContainerId(containerId string) (*domain.VirtualHost, error) {
	var host = &domain.VirtualHost{}

	return host, r.
		db.
		Model(host).
		Where(`container = ?`, containerId).
		Column(`virtual_host.*`).
		Relation(`Pool`).
		Relation(`Client`).
		Select()
}
