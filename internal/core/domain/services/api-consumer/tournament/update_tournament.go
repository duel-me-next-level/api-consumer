package tournament

import "time"

type UpdateTournamentInputDto struct {
	ID      int       `json:"id"`
	BeginAt time.Time `json:"beginAt"`
	EndAt   time.Time `json:"endAt"`
	League  int       `json:"league"`
}

type UpdateTournamentOutputDto struct {
	ID      int
	BeginAt time.Time
	EndAt   time.Time
	League  *League
}

type UpdateTournamentUseCase struct {
	TournamentRepository TournamentRepository
	LeagueRepository     LeagueRepository
}

func NewUpdateTournamentUseCase(tournamentRepository TournamentRepository, leagueRepository LeagueRepository) *UpdateTournamentUseCase {
	return &UpdateTournamentUseCase{TournamentRepository: tournamentRepository, LeagueRepository: leagueRepository}
}

func (u *UpdateTournamentUseCase) Execute(input UpdateTournamentInputDto) (*UpdateTournamentOutputDto, error) {
	league, err := u.LeagueRepository.GetByID(input.League)
	if err != nil {
		return nil, err
	}
	tournament, err := u.TournamentRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	tournament.BeginAt = input.BeginAt
	tournament.EndAt = input.EndAt
	tournament.League = league
	err = u.TournamentRepository.Update
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
