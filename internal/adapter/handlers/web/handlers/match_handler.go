package web

import (
	"net/http"

	"github.com/duel-me-next-level/api-consumer/internal/app/adapter/repository"
	"github.com/duel-me-next-level/api-consumer/internal/app/web/presenter"
)

// MatchDataHandler handles HTTP requests for match data
func MatchDataHandler(w http.ResponseWriter, r *http.Request) {
	// Get the match ID from the request URL
	matchID := r.URL.Query().Get("id")

	// Get the match data from the repository
	repo := repository.Repository{}
	matchData, err := repo.GetMatchData(matchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the presenter to format the match data as JSON
	matchDataJSON, err := presenter.FormatMatchDataAsJSON(matchData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(matchDataJSON)
}
