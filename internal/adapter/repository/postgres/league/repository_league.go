package league

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type leagueRepository struct {
	db *sql.DB
}

func NewLeagueRepository(db *sql.DB) LeagueRepository {
	return &leagueRepository{
		db: db,
	}
}

func (r *leagueRepository) GetByID(ctx context.Context, id int) (*League, error) {
	row := r.db.QueryRowContext(ctx, "SELECT league_id, live_supported, matches FROM leagues WHERE league_id=$1", id)
	var league League
	if err := row.Scan(&league.LeagueID, &league.LiveSupported, &league.Matches); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("league not found")
		}
		return nil, fmt.Errorf("failed to retrieve league: %w", err)
	}
	return &league, nil
}

func (r *leagueRepository) GetAll(ctx context.Context) ([]*League, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT league_id, live_supported, matches FROM leagues")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve leagues: %w", err)
	}
	defer rows.Close()

	leagues := make([]*League, 0)
	for rows.Next() {
		var league League
		if err := rows.Scan(&league.LeagueID, &league.LiveSupported, &league.Matches); err != nil {
			return nil, fmt.Errorf("failed to retrieve league: %w", err)
		}
		leagues = append(leagues, &league)
	}
	return leagues, nil
}

func (r *leagueRepository) Create(ctx context.Context, league *League) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO leagues (league_id, live_supported, matches) VALUES ($1, $2, $3)", league.LeagueID, league.LiveSupported, league.Matches)
	if err != nil {
		return fmt.Errorf("failed to create league: %w", err)
	}
	return nil
}
func (r *leagueRepository) Update(ctx context.Context, league *League) error {
	const query = `UPDATE leagues SET live_supported = $1, matches = $2 WHERE league_id = $3`

	_, err := r.db.ExecContext(ctx, query, league.LiveSupported, league.Matches, league.LeagueID)
	return err
}

func (r *leagueRepository) Delete(ctx context.Context, id int) error {
	const query = `DELETE FROM leagues WHERE league_id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
