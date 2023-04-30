package repository

import (
	"time"

	domain "github.com/duel-me-next-level/api-consumer/internal/core/domain/model"
)

type TournamentService interface {
	GetAll() (domain.Tournament, error)
	Create(BeginAt time.Time, EndAt time.Time, League domain.League) (domain.Tournament, error)
}
