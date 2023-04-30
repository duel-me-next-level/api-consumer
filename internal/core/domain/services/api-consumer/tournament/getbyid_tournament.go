package service

import (
	"github.com/duel-me-next-level/api-consumer/internal/core/domain/model"
	repository "github.com/duel-me-next-level/api-consumer/internal/core/domain/ports"
)

type GetTournamentByIDInputDto struct {
	ID int32 `json:"id"`
}

type GetTournamentByIDOutputDto struct {
	Tournament *model.Tournament
}

type GetTournamentByIDUseCase struct {
	TournamentRepository repository.TournamentRepository
}

func NewGetTournamentByIDUseCase(tournamentRepository repository.TournamentRepository) *GetTournamentByIDUseCase {
	return &GetTournamentByIDUseCase{TournamentRepository: tournamentRepository}
}

func (u *GetTournamentByIDUseCase) Execute(input GetTournamentByIDInputDto) (*GetTournamentByIDOutputDto, error) {
	tournament, err := u.TournamentRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	return &GetTournamentByIDOutputDto{
		Tournament: tournament,
	}, nil
}
