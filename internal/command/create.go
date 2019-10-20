package command

import (
	"errors"

	"github.com/go-pg/pg"
	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
	"go.uber.org/zap"
)

var FreePoolIsAbsent = errors.New(`free pools is absent`)

type create struct {
	pool   domain.PoolRepository
	logger *zap.SugaredLogger
}

func NewCreate(
	pool domain.PoolRepository,
	logger *zap.SugaredLogger,
) *create {
	return &create{
		pool:   pool,
		logger: logger,
	}
}

func (c *create) Exec(req *context.Request) (*Result, error) {
	if _, err := c.pool.GetFreePool(); err != nil {
		if err == pg.ErrNoRows {
			c.logger.Error(`Unable to find any vacant pool`)
			return nil, FreePoolIsAbsent
		}
	}

	return nil, nil
}

func (c *create) Precept() string {
	return `create`
}
