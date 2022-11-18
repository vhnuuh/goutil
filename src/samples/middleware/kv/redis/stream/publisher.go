package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"math/rand"
)

func main() {
	log.Println("Publisher started")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "192.168.136.24", "6379"),
		Password: "password",
	})
	_, err := rdb.Ping().Result()

	if err != nil {
		log.Fatal("Unable to connect to Redis", err)
	}

	log.Println("Connected to Redis server")

	for i := 0; i < 3000; i++ {
		err = publishTicketReceivedEvent(rdb)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func publishTicketReceivedEvent(rdb *redis.Client) error {
	log.Println("Publishing event to Redis")

	err := rdb.XAdd(&redis.XAddArgs{
		Stream:       "tickets",
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values: map[string]interface{}{
			"whatHappened": "ticket received",
			"ticketID":     rand.Intn(100000000),
			"ticketData":   "some ticket data",
		},
	}).Err()

	return err
}
