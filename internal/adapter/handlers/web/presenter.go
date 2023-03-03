package web

import (
	"encoding/json"

	"github.com/duel-me-next-level/api-consumer/internal/app/adapter/model"
)

// FormatMatchDataAsJSON formats match data as JSON
func FormatMatchDataAsJSON(matchData *model.MatchDataDTOOutput) ([]byte, error) {
	matchDataJSON, err := json.Marshal(matchData)
	if err != nil {
		return nil, err
	}

	return matchDataJSON, nil
}
