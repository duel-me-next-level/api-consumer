package postgres

import (
	"github.com/duel-me-next-level/api-consumer/app/adapter/model"
	"github.com/duel-me-next-level/api-consumer/internal/app/adapter/model"
	"github.com/duel-me-next-level/api-consumer/internal/infra/database"
)

// Repository represents the repository for e-sports data
type Repository struct {
	DB *database.DB
}

// SaveMatchData saves match data to the database
func (r *Repository) SaveMatchData(matchData *model.MatchDataDTOInput) error {
	// Prepare the SQL statement
	stmt, err := r.DB.Prepare("INSERT INTO matches (id, match_date, home_team, away_team, home_score, away_score) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(matchData.ID, matchData.MatchDate, matchData.HomeTeam, matchData.AwayTeam, matchData.HomeScore, matchData.AwayScore)
	if err != nil {
		return err
	}

	return nil
}

// GetMatchData retrieves match data from the database
func (r *Repository) GetMatchData(id string) (*model.MatchDataDTOOutput, error) {
	// Prepare the SQL statement
	stmt, err := r.DB.Prepare("SELECT id, match_date, home_team, away_team, home_score, away_score FROM matches WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement
	row := stmt.QueryRow(id)

	// Scan the result into a MatchData struct
	var matchData model.MatchData
	err = row.Scan(&matchData.ID, &matchData.MatchDate, &matchData.HomeTeam, &matchData.AwayTeam, &matchData.HomeScore, &matchData.AwayScore)
	if err != nil {
		return nil, err
	}

	return &matchData, nil
}
