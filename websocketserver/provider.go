package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
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

type WMessage struct {
	// the json tag means this will serialize as a lowercased field
	Message string `json:"message"`
	ch      chan string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Chann = make(chan string)
var client *redis.Client

// var upgrader = websocket.Upgrader{}

func RedisURL() string {
	return "quotetick.njaryy.ng.0001.use2.cache.amazonaws.com:6379"
}

func printit(con *websocket.Conn, pubsub *redis.PubSub) {
	fmt.Println("IT WORKS< CALLING PRINTIT")
	defer con.Close()
	for {

		_, err := pubsub.ReceiveMessage()
		if err != nil {
			log.Println("Channel", err)
			break
		}

		time.Sleep(2 * time.Second)
		err = websocket.WriteJSON(con, interface{}("THIS IS WORKING WITH NEW SETUP"))
		fmt.Println("WORKS -> ", time.Second)
		if err != nil {
			fmt.Println("THERE WAS AN ERROR -> ", err)
			con.Close()
			break
		}
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(<-Chann)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()
	fmt.Println("it was called")

	pubsub := client.Subscribe("")
	defer pubsub.Close()

	go printit(ws, pubsub)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("ERROR ON ECHO -> read: -> ", err)
			// break
			ws.Close()
			break
		}
		fmt.Println(msg)

	}
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

	http.HandleFunc("/echo", echo)
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("THE ERROR ON SPING UP -> ", err)
	}
}
