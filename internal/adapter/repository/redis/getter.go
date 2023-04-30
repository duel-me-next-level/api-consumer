package redis

/* package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func getter() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	id, err := client.Get("api:match:1").Int64()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("match id: ", id)

	// Output: match id: 1
}
*/
