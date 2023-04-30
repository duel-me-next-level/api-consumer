package redis

/* package main

import (
	"fmt"

	"github.com/duel-me-next-level/api-consumer/app/domain/model"
	"github.com/go-redis/redis"
)

func setter() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	match := &match.Match{
		ID:          1,
		Tournament:  "The International",
		Team1:       "Team Liquid",
		Team2:       "Team Secret",
		StartTime:   "2022-08-20T15:00:00Z",
		EndTime:     "2022-08-20T17:00:00Z",
		Winner:      "Team Liquid",
		ScoreTeam1:  2,
		ScoreTeam2:  0,
		MatchStatus: "completed",
	}

	err := client.Set(fmt.Sprintf("api:match:%d", match.ID), match.ID, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("match saved")

	// Output: match saved
}
*/
