package model

type League struct {
	LeagueID      string
	LiveSupported bool
}

func NewLeague(leagueID string, liveSupported bool) *League {
	return &League{
		LeagueID:      leagueID,
		LiveSupported: liveSupported,
	}
}

//struct - Matches       []Match

//[]*Match - funcNewLeague

//return Matches:       matches,
