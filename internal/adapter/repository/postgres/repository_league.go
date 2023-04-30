package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/duel-me-next-level/api-consumer/internal/core/domain/model"
)

type LeagueRepository struct {
	db *sql.DB
}

func NewLeagueRepository(db *sql.DB) *LeagueRepository {
	return &LeagueRepository{
		db: db,
	}
}

func (r *LeagueRepository) GetByID(ctx context.Context, id int) (*model.League, error) {
	row := r.db.QueryRowContext(ctx, "SELECT league_id, live_supported, matches FROM leagues WHERE league_id=$1", id)
	var league model.League
	if err := row.Scan(&league.LeagueID, &league.LiveSupported); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("league not found")
		}
		return nil, fmt.Errorf("failed to retrieve league: %w", err)
	}
	return &league, nil
}

func (r *LeagueRepository) GetAll(ctx context.Context) ([]*model.League, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT league_id, live_supported, matches FROM leagues")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve leagues: %w", err)
	}
	defer rows.Close()

	leagues := make([]*model.League, 0)
	for rows.Next() {
		var league model.League
		if err := rows.Scan(&league.LeagueID, &league.LiveSupported); err != nil {
			return nil, fmt.Errorf("failed to retrieve league: %w", err)
		}
		leagues = append(leagues, &league)
	}
	return leagues, nil
}

func (r *LeagueRepository) Create(ctx context.Context, league *model.League) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO leagues (league_id, live_supported, matches) VALUES ($1, $2, $3)", league.LeagueID, league.LiveSupported)
	if err != nil {
		return fmt.Errorf("failed to create league: %w", err)
	}
	return nil
}
func (r *LeagueRepository) Update(ctx context.Context, league *model.League) error {
	const query = `UPDATE leagues SET live_supported = $1, matches = $2 WHERE league_id = $3`

	_, err := r.db.ExecContext(ctx, query, league.LiveSupported, league.LeagueID)
	return err
}

func (r *LeagueRepository) Delete(ctx context.Context, id int) error {
	const query = `DELETE FROM leagues WHERE league_id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
