package web

import (
	"encoding/json"

	service "github.com/duel-me-next-level/api-consumer/internal/core/domain/services/api-consumer/tournament"
)

// FormatMatchDataAsJSON formats match data as JSON
func FormatMatchDataAsJSON(matchData *service.CreateTournamentOutputDto) ([]byte, error) {
	matchDataJSON, err := json.Marshal(matchData)
	if err != nil {
		return nil, err
	}

	return matchDataJSON, nil
}
