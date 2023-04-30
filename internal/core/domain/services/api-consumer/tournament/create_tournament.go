package service

import (
	"time"

	"github.com/duel-me-next-level/api-consumer/internal/core/domain/model"
	repository "github.com/duel-me-next-level/api-consumer/internal/core/domain/ports"
)

type CreateTournamentInputDto struct {
	ID            int32         `json:"id"`
	BeginAt       time.Time     `json:"begin_at"`
	EndAt         time.Time     `json:"end_at"`
	League        *model.League `json:"league"`
	VideoGameID   int32         `json:"video_game_id"`
	videoGameName string        `json:"video_game_name"`
	winnerID      int32         `json:"winner_id"`
	winnerType    string        `json:"winner_type"`
}

type CreateTournamentOutputDto struct {
	ID      int32
	BeginAt time.Time
	EndAt   time.Time
	League  model.League
}

type CreateTournamentUseCase struct {
	TournamentRepository repository.TournamentRepository
	LeagueRepository     repository.LeagueRepository
}

func NewCreateTournamentUseCase(tournamentRepository repository.TournamentRepository, leagueRepository repository.LeagueRepository) *CreateTournamentUseCase {
	return &CreateTournamentUseCase{TournamentRepository: tournamentRepository, LeagueRepository: leagueRepository}
}

func (u *CreateTournamentUseCase) Execute(input CreateTournamentInputDto) (*CreateTournamentOutputDto, error) {

	tournament := model.NewTournament(input.BeginAt, input.EndAt, input.ID, *input.League, input.VideoGameID, input.winnerID, input.winnerType)
	err := u.TournamentRepository.Create(tournament)
	if err != nil {
		return nil, err
	}

	return &CreateTournamentOutputDto{
		ID:      tournament.ID,
		BeginAt: tournament.BeginAt,
		EndAt:   tournament.EndAt,
		League:  *tournament.League,
	}, nil
}
