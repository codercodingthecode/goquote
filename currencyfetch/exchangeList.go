package currencyfetch

// List of rest api's to fetch currency pairs list
func exchangeList() map[string]string {
	urls := map[string]string{
		// "hitbtc": "https://api.hitbtc.com/api/2/public/symbol", // response -> id -> "ETHUSD" **
		"gdax": "https://api.gdax.com/products", // response -> id -> "eth" **
		// "binance":  "https://api.binance.com/api/v1/exchangeInfo",          // it needs to dig into json, response -> symbol -> baseAsset -> "ETH" **
		// "poloniex": "https://poloniex.com/public?command=returnCurrencies", // response -> "ETH"   ; the crypto is in the key **
		// https://poloniex.com/public?command=returnTicker  *change to this end point
		// "bitmex": "https://www.bitmex.com/api/v1/stats", // needs to look into it further to find pairs                 // response -> rootSymbol -> "ETH" **
		// "bitfinex": "https://api.bitfinex.com/v1/symbols", // response -> "ethbtc"  || it takes a pair rather than the currency symbol, needs discussion on it
		// "bittrex": "https://bittrex.com/api/v1.1/public/getcurrencies", // response -> result -> Currency -> "ETH" || it uses rest api already.
	}

	return urls
}
