package exchanges

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// _ "github.com/go-sql-driver/mysql"
	ws "github.com/gorilla/websocket"
)

// CoinRest interface holder for rest api incoming data
type CoinRest interface{}

// GDAXI implements funcs
type GDAXI interface {
	fetchGdax()
	restAPICaller()
	SubscribeToGDAX()
}

var gdaxRestURL = "https://api.gdax.com/products" // response -> id -> "eth" **
// func main() {
// 	errCh := make(chan error)
// 	currencies, err := entity.FetchRowsFromDB("gdax")
// 	if err == nil {
// 		go entity.ProcessQuoteTickUpdates()
// 		go subscribeToGDAX(currencies)

// 		select {
// 		case err := <-errCh:
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 		}
// 	}
// }

// Fetchs the currency list from GDAX thru rest call which had the same data structure
// Function: Receives two arguemts and feed data thru the channel back to the caller:
// - url string which is the exchange rest api
// - c channel which is where it feeds the data back to the caller once ready
// it call "restAPICaller" function passing the rest url and a coin interface pointer
// to retrieve the data
func fetchGdax(url string) []string {
	cList := make([]string, 0)
	var coins CoinRest

	restAPICaller(url, &coins)

	for _, coin := range coins.([]interface{}) {
		s := coin.(map[string]interface{})
		cList = append(cList, s["id"].(string))
	}
	// fmt.Println("GDAX REST -> ", cList)
	// c <- cList
	fmt.Println("this is a list of pairs --> ", cList)
	return cList
}

// Helper function common used to by the function below to call the rest api for each exchange.
// Function: Receives two arguments and returns the data by pointer to the caller:
// - url string which is the exchange rest api
// - coins pointer interface which receives the data from the rest call and returns to the caller function.
func restAPICaller(url string, coins *CoinRest) {
	fmt.Println("THIS IS INSIDE RESTAPICALLER IN GDAX")
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("there was an error -> ", err)
	}
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(bytes, &coins)
}

func SubscribeToGDAX(chStatus <-chan string) {

	var writeData interface{}
	var tickRawData interface{}
	var websocket ws.Dialer

	// fmt.Println("WEBSOCKETLIST --> ", WebsocketList()[k])

	wsConn, _, wsError := websocket.Dial("wss://ws-feed.gdax.com", nil) // this will have to change per loop or concurrency
	if wsError != nil {
		fmt.Println("There was an error dialing in -> ", wsError)
	}

	b, _ := json.Marshal(fetchGdax(gdaxRestURL))

	bigString := fmt.Sprintf(`{"type":"subscribe","channels": [{"name": "ticker", "product_ids": %s}]}`, b)

	// fmt.Println("FROM SOCKECT STRING BUILD -> ", bigString)

	// fmt.Println("From exchangeString -> ", testingString)

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

		tickData := tickRawData.(map[string]interface{})

		fmt.Println("This is the feed -> ", tickData)
		// ch <- tickRawData
		// cb(<-ch)
		// exchanges.Gdax(tickRawData)
	}

	// var wsDialer ws.Dialer
	// // loc, _ := time.LoadLocation("UTC")
	// fmt.Println("RUNNING SUB GDAX")

	// wsConn, _, err := wsDialer.Dial("wss://ws-feed.gdax.com", nil) // hoist this to a variable above for code clarity
	// if err != nil {
	// 	fmt.Println("THERE WAS AN ERROR DIALING IN")
	// 	println(err.Error())
	// }

	// str := fmt.Sprintf(`{"type":"subscribe","channels": [{"name": "ticker", "product_ids": [%s]}]}`, fetchGdax(gdaxRestURL))
	// res := map[string]interface{}{}
	// json.Unmarshal([]byte(str), &res)

	// if err := wsConn.WriteJSON(res); err != nil {
	// 	fmt.Println("THERE WAS AN ERROR WRITING JSON --> ", err.Error())
	// }

	// // tick := gdax.Message{}
	// var message interface{}
	// // var marketCurrencyLastPrice = make(map[string]float64)

	// for {
	// 	// wsConn.SetReadDeadline(time.Now().Add(10 * time.Second))
	// 	fmt.Println("THIS IS INSIDE FOR LOOP FOR WEBSOCKET")
	// 	if err1 := wsConn.ReadJSON(&message); err != nil {
	// 		fmt.Println("ERROR READING FROM JSON --> ", err1)
	// 		panic(err1.Error())
	// 	}
	// 	fmt.Println("THIS IS INCOME DATA MESSAGE -> ", message)
	// 	// tick := message.(map[string]interface{})

	// 	// if tick["high_24h"] == nil {
	// 	// 	continue
	// 	// }
	// 	// last, _ := strconv.ParseFloat(tick["price"].(string), 64)
	// 	// marketName, _ := tick["product_id"].(string)

	// 	// prevLast, ok := marketCurrencyLastPrice[marketName]
	// 	// if ok {
	// 	// 	if prevLast == last {
	// 	// 		continue
	// 	// 	}
	// 	// }
	// 	// marketCurrencyLastPrice[marketName] = last

	// 	// high, _ := strconv.ParseFloat(tick["high_24h"].(string), 64)
	// 	// low, _ := strconv.ParseFloat(tick["low_24h"].(string), 64)
	// 	// open, _ := strconv.ParseFloat(tick["open_24h"].(string), 64)
	// 	// volume, _ := strconv.ParseFloat(tick["volume_24h"].(string), 64)
	// 	// change := last - open
	// 	// changePercent := change / open

	// 	// currencySymbol := entity.MarketNameToCurrencySymbolMap[marketName]
	// 	// baseSymbol := entity.MarketNameToBaseNameMap[marketName]

	// 	// val := entity.QuoteTick{baseSymbol, currencySymbol, "gdax-" + currencySymbol + "-" + baseSymbol, open, high, low, last, volume, time.Now().In(loc), change, changePercent}
	// 	// entity.QuoteUpdates <- val
	// 	fmt.Println("THIS IS IN THE LOOP FOR SOCKET INFO --------->>>>> ")
	// }
}
