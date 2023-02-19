package tournament_interface

import (
	"time"

	domain "github.com/duel-me-next-level/api-consumer/internal/core/domain/model/tournament"
)

type TournamentService interface {
	GetAll() (domain.Tournament, error)
	Create(BeginAt time.Time, EndAt time.Time, LeagueID int) (domain.Tournament, error)
}
