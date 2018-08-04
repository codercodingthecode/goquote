package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	// _ "github.com/go-sql-driver/mysql"
)

type quoteTicks struct {
	BaseName       string
	CurrencySymbol string
	MarketName     string
	Open           float64
	High           float64
	Low            float64
	Last           float64
	Volume         float64
	TimeStamp      time.Time
	Change         float64
	ChangePercent  float64
}

func RedisURL() string {
	return "quotetick.njaryy.ng.0001.use2.cache.amazonaws.com:6379"
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     RedisURL(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer client.Close()
	pong, err := client.Ping().Result()
	fmt.Println("from redis --------------------->>>>   ", pong, err)

	 pubsub := client.Subscribe("QuoteTicks")
	 defer pubsub.Close()

	 for {
	 	msg, err := pubsub.ReceiveMessage()
	 	if err != nil {
	 		fmt.Println("Erros ->>  ", err)
	 		break
	 	}
	 	 //var quoteTick quoteTicks
	 	fmt.Println(msg)
	 	 //err1 := quoteTick.UnmarshalBinary([]byte(msg.Payload))
	 	 //if err1 == nil {
	 	// 	marketsForCandle <- quoteTick
	 	 //} else {
	 	// 	fmt.Println(err1)
	 	 //}
	 }
}
