package socketsclient

import (
	"encoding/json"
	"fmt"

	ws "github.com/gorilla/websocket"
)

// const gdax = "wss://ws-feed.gdax.com"

// WSHitbtc is pulling data
func WSHitbtc(cL map[string][]string, ch chan interface{}) {

	// fmt.Println("DATA COMING FROM MAIN ---------------> ", cL["gdax"])

	// @TODO:
	// move this code to socket.go
	var writeData interface{}
	var websocket ws.Dialer
	wsConn, _, wsError := websocket.Dial(WebsocketList()["gdax"], nil) // this will have to change per loop or concurrency
	if wsError != nil {
		fmt.Println("There was an error dialing in -> ", wsError)
	}

	// fmt.Println("This is wsConn -> ", wsConn.)
	// this will have to change to either for loop or change
	// on each call per concurrent exchange
	b, _ := json.Marshal(cL["gdax"])

	// @TODO:
	// this will change based on exchange, move to a proper go file to be setup
	bigString := fmt.Sprintf(`{"type":"subscribe","channels": [{"name": "ticker", "product_ids": %s}]}`, b)

	// @TODO:
	// move this code to socket.go
	writeDataError := json.Unmarshal([]byte(bigString), &writeData)
	if writeDataError != nil {
		fmt.Println("There was an eror unmarshalling ", writeDataError)
	}

	// fmt.Println("writedata json ", writeData)

	writeJSONError := wsConn.WriteJSON(writeData)
	if writeJSONError != nil {
		fmt.Println("there was an error writing json -> ", writeJSONError)
	}

	var tickRawData interface{}
	for {
		// break
		readError := wsConn.ReadJSON(&tickRawData)

		if readError != nil {
			fmt.Println("Error in pulling in data -> ", readError)
		}

		// tickData := tickRawData.(map[string]interface{})

		// fmt.Println("This is the feed -> ", tickData)
		ch <- tickRawData
	}

}
