package tournament

import "time"

type Tournament struct {
	BeginAt time.Time
	EndAt   time.Time
	ID      int
	League  *League
}

func NewTournament(beginAt time.Time, endAt time.Time, id string, league *League) *Tournament {
	return &Tournament{
		BeginAt: beginAt,
		EndAt:   endAt,
		ID:      id,
		League:  league,
	}
}
