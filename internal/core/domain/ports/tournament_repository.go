package repository

import domain "github.com/duel-me-next-level/api-consumer/internal/core/domain/model"

type TournamentRepository interface {
	GetByID(id int32) (*domain.Tournament, error)
	Create(tournament *domain.Tournament) error
	Update(tournament *domain.Tournament) error
	FindAll() ([]*domain.Tournament, error)
	Delete(id int32) error
}
