package repository

import "context"

type LeagueRepository interface {
	GetByID(ctx context.Context, id int) (*League, error)
	GetAll(ctx context.Context) ([]*League, error)
	Create(ctx context.Context, league *League) error
	Update(ctx context.Context, league *League) error
	Delete(ctx context.Context, id int) error
}
