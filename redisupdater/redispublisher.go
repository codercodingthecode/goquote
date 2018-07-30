package redispublisher

import (
	"sort"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var QuoteUpdates = make(chan QuoteTick, 0)

func RedisURL() string {
	// return "127.0.0.1:6379"
	return "quotetick.fcryea.ng.0001.use1.cache.amazonaws.com:6379"
	// return "uptick-stream-redis.njaryy.0001.use2.cache.amazonaws.com:6379"
}

func ProcessQuoteTickUpdates() {
	client := redis.NewClient(&redis.Options{
		Addr: RedisURL(),
	})
	defer client.Close()

	for val := range QuoteUpdates {
		content, _ := val.MarshalBinary()
		_, err := client.Publish("QuoteTicks", content).Result()
		if err != nil {
			panic(err)
		}
		client.Publish("QuoteTicks-"+val.MarketName, content).Result()
		client.Set("QuoteTicks-"+val.MarketName, content, 0).Result()
	}
}

var ExchangeSpecificOrderBook = make(map[string]OrderBook)
var ExchangeSpecificOrderBookMutex = &sync.Mutex{}

func DoEvery(d time.Duration, f func(time.Time, *redis.Client)) {
	client := redis.NewClient(&redis.Options{
		Addr: RedisURL(),
	})
	defer client.Close()
	for x := range time.Tick(d) {
		go f(x, client)
	}
}

func ProcessOrderBookUpdates(t time.Time, client *redis.Client) {
	ExchangeSpecificOrderBookMutex.Lock()
	for _, currVal := range ExchangeSpecificOrderBook {
		orderBookToSend := OrderBookToSend{currVal.BaseName, currVal.CurrencySymbol, currVal.MarketName, make([]float64, 0), make([]float64, 0), make([]float64, 0), make([]float64, 0)}
		var asks []float64
		for k := range currVal.AskMap {
			asks = append(asks, k)
		}
		sort.Float64s(asks)

		var bids []float64
		for k := range currVal.BidMap {
			bids = append(bids, k)
		}
		sort.Float64s(bids)

		if len(bids) > 200 {
			bids = bids[len(bids)-200 : len(bids)-1]
		}

		for _, k := range bids {
			// if k < lowerLimit {
			// 	continue
			// }
			orderBookToSend.Bids = append(orderBookToSend.Bids, k)
			orderBookToSend.BidsQuantity = append(orderBookToSend.BidsQuantity, currVal.BidMap[k])
		}

		if len(asks) > 200 {
			asks = asks[0:200]
		}

		for _, k := range asks {
			// if k > upperLimit {
			// 	break
			// }
			orderBookToSend.Asks = append(orderBookToSend.Asks, k)
			orderBookToSend.AsksQuantity = append(orderBookToSend.AsksQuantity, currVal.AskMap[k])
		}
		content, _ := orderBookToSend.MarshalBinary()

		_, err := client.Publish("OrderBooks-"+currVal.MarketName, content).Result()
		if err != nil {
			panic(err)
		}

		client.Set("OrderBooks-"+currVal.MarketName, content, 0).Result()
	}
	ExchangeSpecificOrderBookMutex.Unlock()
}
