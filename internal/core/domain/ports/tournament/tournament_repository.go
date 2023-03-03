package tournament_interface

import domain "github.com/duel-me-next-level/api-consumer/internal/core/domain/model/tournament"

type TournamentRepository interface {
	GetByID(id int) (*domain.Tournament, error)
	Create(tournament *domain.Tournament) error
	Update(tournament *domain.Tournament) error
	FindAll() ([]*domain.Tournament, error)
	Delete(id int) error
}
