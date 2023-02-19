type FindAllTournamentsInputDto struct {
}

type FindAllTournamentsOutputDto struct {
	Tournaments []Tournament
}

type FindAllTournamentsUseCase struct {
	TournamentRepository TournamentRepository
}

func NewFindAllTournamentsUseCase(tournamentRepository TournamentRepository) *FindAllTournamentsUseCase {
	return &FindAllTournamentsUseCase{TournamentRepository: tournamentRepository}
}

func (u *FindAllTournamentsUseCase) Execute(input FindAllTournamentsInputDto) (*FindAllTournamentsOutputDto, error) {
	tournaments, err := u.TournamentRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &FindAllTournamentsOutputDto{
		Tournaments: tournaments,
	}, nil
}
