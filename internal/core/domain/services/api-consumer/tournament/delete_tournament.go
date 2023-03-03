package tournament

type DeleteTournamentInputDto struct {
	ID int
}

type DeleteTournamentOutputDto struct {
}

type DeleteTournamentUseCase struct {
	TournamentRepository domain.TournamentRepository
}

func NewDeleteTournamentUseCase(tournamentRepository domain.TournamentRepository) *DeleteTournamentUseCase {
	return &DeleteTournamentUseCase{TournamentRepository: tournamentRepository}
}

func (u *DeleteTournamentUseCase) Execute(input DeleteTournamentInputDto) (*DeleteTournamentOutputDto, error) {
	err := u.TournamentRepository.Delete(input.ID)
	if err != nil {
		return nil, err
	}
	return &DeleteTournamentOutputDto{}, nil
}
