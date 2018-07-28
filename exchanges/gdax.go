package exchanges

import (
	"fmt"
)

// Gdax formats the data from raw websocket
func Gdax(liveRawData chan interface{}) {
	fmt.Println("GDAX LIVE DATA")
	for {
		fmt.Println("GDAX -> ", <-liveRawData)
	}
}
