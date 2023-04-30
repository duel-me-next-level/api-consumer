package repository

import (
	"context"

	"github.com/duel-me-next-level/api-consumer/internal/core/domain/model"
)

type LeagueRepository interface {
	GetByID(ctx context.Context, id int) (*model.League, error)
	GetAll(ctx context.Context) ([]*model.League, error)
	Create(ctx context.Context, league *model.League) error
	Update(ctx context.Context, league *model.League) error
	Delete(ctx context.Context, id int) error
}
