package socketsclient

import (
	"encoding/json"
	"fmt"

	ws "github.com/gorilla/websocket"
	// "github.com/codercodingthecode/daix/currencyFetch/exchanges"
)

// SocketStream common function for all exchange, returns raw data
// func SocketStream(cL map[string][]string, ch chan interface{}) {
func SocketStream(exchange string, pairs []string, ch chan interface{}) {
	var writeData interface{}
	var tickRawData interface{}
	var websocket ws.Dialer

	// fmt.Println("WEBSOCKETLIST --> ", WebsocketList()[k])

	wsConn, _, wsError := websocket.Dial(WebsocketList()[exchange], nil) // this will have to change per loop or concurrency
	if wsError != nil {
		fmt.Println("There was an error dialing in -> ", wsError)
	}

	b, _ := json.Marshal(pairs)

	// move this to struct
	bigString := fmt.Sprintf(`{"type":"subscribe","channels": [{"name": "ticker", "product_ids": %s}]}`, b)

	writeDataError := json.Unmarshal([]byte(bigString), &writeData)
	if writeDataError != nil {
		fmt.Println("There was an eror unmarshalling ", writeDataError)
	}
	writeJSONError := wsConn.WriteJSON(writeData)
	if writeJSONError != nil {
		fmt.Println("there was an error writing json -> ", writeJSONError)
	}
	for {
		readError := wsConn.ReadJSON(&tickRawData)
		if readError != nil {
			fmt.Println("Error in pulling in data -> ", readError)
		}

		// tickData := tickRawData.(map[string]interface{})

		// fmt.Println("This is the feed -> ", tickData)
		ch <- tickRawData
		// cb(<-ch)
		// exchanges.Gdax(tickRawData)
	}

}
