package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	redisServer string = "127.0.0.1"
	port        string = "6379"
	password    string = ""
)

func redisConnect() (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisServer + ":" + port,
		Password: password,
		DB:       0, // use default DB
	})

	return
}

func pubMessage(channel, msg string) {
	rdb := redisConnect()
	rdb.Publish(context.Background(), channel, msg)
}

func subMessage(channel string) {
	rdb := redisConnect()
	pubsub := rdb.Subscribe(context.Background(), channel)
	_, err := pubsub.Receive(context.Background())
	if err != nil {
		panic(err)
	}

	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func Test() {
	fmt.Println("enter redis test")
	channel := "hello"
	//sub
	go func() {
		subMessage(channel)
	}()
	//pub
	go func() {
		ticker := time.NewTicker(time.Second)
		nums := 0
		for {
			select {
			case <-ticker.C:
				fmt.Println("goto pub msg")
				pubMessage(channel, fmt.Sprintf("hello %d times", nums))
				nums += 1
			}
		}
	}()
}
