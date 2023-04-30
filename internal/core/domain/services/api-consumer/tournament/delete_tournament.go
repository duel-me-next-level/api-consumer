package service

import repository "github.com/duel-me-next-level/api-consumer/internal/core/domain/ports"

type DeleteTournamentInputDto struct {
	ID int32
}

type DeleteTournamentUseCase struct {
	TournamentRepository repository.TournamentRepository
}

func NewDeleteTournamentUseCase(tournamentRepository repository.TournamentRepository) *DeleteTournamentUseCase {
	return &DeleteTournamentUseCase{TournamentRepository: tournamentRepository}
}

func (u *DeleteTournamentUseCase) Execute(input DeleteTournamentInputDto) error {
	err := u.TournamentRepository.Delete(input.ID)
	if err != nil {
		return err
	}
	return nil
}
