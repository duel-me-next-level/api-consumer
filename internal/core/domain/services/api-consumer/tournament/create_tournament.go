package tournament

import "time"

type CreateTournamentInputDto struct {
	BeginAt  time.Time `json:"begin_at"`
	EndAt    time.Time `json:"end_at"`
	LeagueID int       `json:"league_id"`
}

type CreateTournamentOutputDto struct {
	ID       int
	BeginAt  time.Time
	EndAt    time.Time
	LeagueID int
}

type CreateTournamentUseCase struct {
	TournamentRepository repository.TournamentRepository
	LeagueRepository     repository.LeagueRepository
}

func NewCreateTournamentUseCase(tournamentRepository domain.TournamentRepository, leagueRepository domain.LeagueRepository) *CreateTournamentUseCase {
	return &CreateTournamentUseCase{TournamentRepository: tournamentRepository, LeagueRepository: leagueRepository}
}

func (u *CreateTournamentUseCase) Execute(input CreateTournamentInputDto) (*CreateTournamentOutputDto, error) {
	league, err := u.LeagueRepository.GetByID(input.LeagueID)
	if err != nil {
		return nil, err
	}

	tournament := domain.NewTournament(input.BeginAt, input.EndAt, league)
	err = u.TournamentRepository.Create(tournament)
	if err != nil {
		return nil, err
	}

	return &CreateTournamentOutputDto{
		ID:       tournament.ID,
		BeginAt:  tournament.BeginAt,
		EndAt:    tournament.EndAt,
		LeagueID: tournament.League.ID,
	}, nil
}
