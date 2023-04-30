package service

import (
	"github.com/duel-me-next-level/api-consumer/internal/core/domain/model"
	repository "github.com/duel-me-next-level/api-consumer/internal/core/domain/ports"
)

type FindAllTournamentsOutputDto struct {
	Tournaments []model.Tournament
}

type FindAllTournamentsUseCase struct {
	TournamentRepository repository.TournamentRepository
}

func NewFindAllTournamentsUseCase(tournamentRepository repository.TournamentRepository) *FindAllTournamentsUseCase {
	return &FindAllTournamentsUseCase{TournamentRepository: tournamentRepository}
}

func (u *FindAllTournamentsUseCase) Execute() (*FindAllTournamentsOutputDto, error) {
	tournaments, err := u.TournamentRepository.FindAll()
	if err != nil {
		return nil, err
	}
	tournamentDto := FindAllTournamentsOutputDto{
		Tournaments: make([]model.Tournament, 0),
	}
	for _, t := range tournaments {
		tournamentDto.Tournaments = append(tournamentDto.Tournaments, *t)
	}

	return &tournamentDto, nil
}
