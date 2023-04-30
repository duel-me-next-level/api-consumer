package model

import "time"

type Match struct {
	BeginAt       time.Time
	DetailedStats map[string]interface{}
	Draw          bool
	EndAt         time.Time
	Forfeit       bool
	GameAdvantage string
	ID            int
}

func NewMatch(beginAt time.Time, detailedStats map[string]interface{}, draw bool, endAt time.Time, forfeit bool, gameAdvantage string, id int) *Match {
	return &Match{
		BeginAt:       beginAt,
		DetailedStats: detailedStats,
		Draw:          draw,
		EndAt:         endAt,
		Forfeit:       forfeit,
		GameAdvantage: gameAdvantage,
		ID:            id,
	}
}
