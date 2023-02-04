package tournament

type GetTournamentByIDInputDto struct {
	ID int `json:"id"`
}

type GetTournamentByIDOutputDto struct {
	Tournament *Tournament
}

type GetTournamentByIDUseCase struct {
	TournamentRepository TournamentRepository
}

func NewGetTournamentByIDUseCase(tournamentRepository TournamentRepository) *GetTournamentByIDUseCase {
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
