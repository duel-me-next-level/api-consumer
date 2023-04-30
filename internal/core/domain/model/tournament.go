package model

import (
	"errors"
	"time"

	"github.com/duel-me-next-level/api-consumer/internal/core/domain/model"
)

type Tournament struct {
	BeginAt       time.Time
	EndAt         time.Time
	ID            int32
	League        *model.League
	VideoGameID   int32
	videoGameName string
	winnerID      int32
	winnerType    string
}

func NewTournament(beginAt time.Time, endAt time.Time, id int32, league domain.League, videoGameID int32, winnerId int32, winnerType string) *Tournament {

	videoGameName, err := getVideoGameName(videoGameID)
	if err != nil {
		panic(err)
	}

	return &Tournament{
		BeginAt:       beginAt,
		EndAt:         endAt,
		ID:            id,
		League:        domain.NewLeague(league.LeagueID, league.LiveSupported),
		VideoGameID:   videoGameID,
		videoGameName: videoGameName,
		winnerID:      winnerId,
		winnerType:    winnerType,
	}
}

func getVideoGameName(videoGameID int32) (string, error) {
	videoGameMap := map[int32]string{
		1: "LoL",
		3: "CS:GO",
		4: "Dota 2",
	}
	videoGameName, err := videoGameMap[videoGameID]
	if !err {
		return "", errors.New("video game not found")
	}
	return videoGameName, nil
}
