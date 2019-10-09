package command

import (
	"chat/src/context"
	"chat/src/domain"
	"errors"
	"github.com/go-pg/pg"
	"go.uber.org/zap"
)

var FreePoolIsAbsent = errors.New(`free pools is absent`)

type create struct {
	precept string
	pool    domain.PoolRepository
	logger  *zap.SugaredLogger
}

func NewCreate(
	precept string,
	pool domain.PoolRepository,
	logger *zap.SugaredLogger,
) *create {
	return &create{
		precept: precept,
		pool:    pool,
		logger:  logger,
	}
}

func (c *create) Exec(req *context.Request) error {
	if _, err := c.pool.GetFreePool(); err != nil {
		if err == pg.ErrNoRows {
			c.logger.Error(`Unable to find any vacant pool`)
			return FreePoolIsAbsent
		}
	}

	return nil
}

func (c *create) Precept() string {
	return c.precept
}
