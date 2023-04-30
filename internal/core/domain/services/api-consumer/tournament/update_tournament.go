package service

import (
	"time"

	"github.com/duel-me-next-level/api-consumer/internal/core/domain/model"
	repository "github.com/duel-me-next-level/api-consumer/internal/core/domain/ports"
)

type UpdateTournamentInputDto struct {
	ID      int32         `json:"id"`
	BeginAt time.Time     `json:"beginAt"`
	EndAt   time.Time     `json:"endAt"`
	League  *model.League `json:"league"`
}

type UpdateTournamentOutputDto struct {
	ID      int32
	BeginAt time.Time
	EndAt   time.Time
	League  *model.League
}

type UpdateTournamentUseCase struct {
	TournamentRepository repository.TournamentRepository
	LeagueRepository     repository.LeagueRepository
}

func NewUpdateTournamentUseCase(tournamentRepository repository.TournamentRepository, leagueRepository repository.LeagueRepository) *UpdateTournamentUseCase {
	return &UpdateTournamentUseCase{TournamentRepository: tournamentRepository, LeagueRepository: leagueRepository}
}

func (u *UpdateTournamentUseCase) Execute(input UpdateTournamentInputDto) (*UpdateTournamentOutputDto, error) {
	tournament, err := u.TournamentRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	tournament.BeginAt = input.BeginAt
	tournament.EndAt = input.EndAt
	tournament.League = input.League
	err = u.TournamentRepository.Update(tournament)
	if err != nil {
		return nil, err
	}

	return &UpdateTournamentOutputDto{
		ID:      tournament.ID,
		BeginAt: tournament.BeginAt,
		EndAt:   tournament.EndAt,
		League:  tournament.League,
	}, nil
}
