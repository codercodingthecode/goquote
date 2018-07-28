package main

import (
	"fmt"

	"github.com/codercodingthecode/goquote/currencyfetch"
	"github.com/codercodingthecode/goquote/socketsclient"
	// "github.com/codercodingthecode/goquote/currencyfetch/socketsclient"
	// "github.com/codercodingthecode/goquote/currencyfetch/socketsclient"
	// "currencyFetch/socketsclient"
	// "github.com/codercodingthecode/goquote/currencyfetch/currencyfetch"
	// "currencyFetch/currencyfetch"
)

func main() {
	ch := make(chan interface{})
	// chOut := make(chan interface{})
	cL := currencyfetch.FetchCurrencies()

	for k, v := range cL {
		fmt.Println("RUNNING ON MAIN")
		go socketsclient.SocketStream(k, v, ch)
	}

	for {
		fmt.Println(<-ch)
	}
	// exchanges.Gdax(ch)

}
