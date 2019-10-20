package repository

import (
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
