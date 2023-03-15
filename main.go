package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const channelName = "exampleChannel"

func main() {
	rdb := openConnection()

	pubsub := rdb.Subscribe(context.Background(), channelName)
	defer pubsub.Close()

	go func() {
		receiverMessage(pubsub)
	}()

	err := rdb.Publish(context.Background(), channelName, "PubSub Example Message").Err()
	if err != nil {
		fmt.Printf("There was an error in publish. Error: %s", err.Error())
	} else {
		fmt.Print("Message published.")
	}
}

func openConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return rdb
}

func receiverMessage(pubsub *redis.PubSub) {
	message, err := pubsub.ReceiveMessage(context.Background())
	if err != nil {
		fmt.Printf("There was an error in receive. Error: %s", err.Error())
	} else {
		fmt.Printf("Message received. Message: %s", message.Payload)
	}
}
