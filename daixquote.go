package main

import (
	"fmt"

	"github.com/codercodingthecode/goquote/exchanges"
)

func main() {

	chStatus := make(chan string)

	go exchanges.SubscribeToGDAX(chStatus)

	fmt.Println("THERE WAS AN ERROR ---> ", <-chStatus)

	// ch := make(chan interface{})
	// chOut := make(chan interface{})
	// cL := currencyfetch.FetchCurrencies()

	// fmt.Println(cL, ch)
	// for exchange, pairs := range cL {
	// 	fmt.Println("RUNNING ON MAIN")
	// 	// dispatch a go routine per exchange on the list
	// 	go socketsclient.SocketStream(exchange, pairs, ch, currencyfetch.ExchangeString)
	// }

	// for {
	// 	// fmt.Println(<-ch)
	// }
	// exchanges.Gdax(ch)

}
