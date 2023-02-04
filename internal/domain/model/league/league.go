package league

type League struct {
	LeagueID      string
	LiveSupported bool
	Matches       []Match
}

func NewLeague(leagueID string, liveSupported bool, matches []*Match) *League {
	return &League{
		LeagueID:      leagueID,
		LiveSupported: liveSupported,
		Matches:       matches,
	}
}
