package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/rs/xid"
	"log"
)

func main() {
	log.Println("Consumer started")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "192.168.136.24", "6379"),
		Password: "password",
	})
	_, err := rdb.Ping().Result()

	if err != nil {
		log.Fatal("Unable to connect to Redis", err)
	}

	log.Println("Connected to Redis server")

	subject := "tickets"
	consumersGroup := "tickets-consumer-group"

	err = rdb.XGroupCreate(subject, consumersGroup, "0").Err()
	if err != nil {
		log.Println(err)
	}

	uniqueID := xid.New().String()

	for {
		entries, err := rdb.XReadGroup(&redis.XReadGroupArgs{
			Group:    consumersGroup,
			Consumer: uniqueID,
			Streams:  []string{subject, ">"},
			Count:    2,
			Block:    0, // 0: block forever
			NoAck:    false,
		}).Result()
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(entries[0].Messages); i++ {
			messageID := entries[0].Messages[i].ID
			values := entries[0].Messages[i].Values
			eventDescription := fmt.Sprintf("%v", values["whatHappened"])
			ticketID := fmt.Sprintf("%v", values["ticketID"])
			ticketData := fmt.Sprintf("%v", values["ticketData"])

			if eventDescription == "ticket received" {
				err := handleNewTicket(ticketID, ticketData)
				if err != nil {
					log.Fatal(err)
				}
				rdb.XAck(subject, consumersGroup, messageID)
			}
		}
	}
}

func handleNewTicket(ticketID string, ticketData string) error {
	log.Printf("Handling new ticket id : %s data %s\n", ticketID, ticketData)
	// time.Sleep(100 * time.Millisecond)
	return nil
}
