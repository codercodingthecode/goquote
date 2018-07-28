package dataoutbound

// XCHGLiveFeed Per exchange final data structure to provide data to front end
// Example: GDAX[...], COINEX[...]
// ** Needs to jsonfy the data fields
type XCHGLiveFeed struct {
	FeedData feedData
}

type feedData struct {
	marketname string // confirm this

	// this two below, in case of BTC-USD, what do they relate to?
	currencySymbol string
	baseSymbol     string

	last          float64 // confirm this, [perhaps instead last, "price" would be more appropriate?]
	high          float64
	low           float64
	open          float64
	volume        float64
	change        float64
	changePercent int
}

// IM TRYING TO GET THE LIVE FEED DATA STRUCTURE OUT OF THE EXAMPLE BELOW TO FEED INTO REDIS/CASSANDRA/FRONTEND

// tick := message.(map[string]interface{})

// if tick["high_24h"] == nil {
// 	continue
// }
// last, _ := strconv.ParseFloat(tick["price"].(string), 64)
// marketName, _ := tick["product_id"].(string)

// prevLast, ok := marketCurrencyLastPrice[marketName]
// if ok {
// 	if prevLast == last {
// 		continue
// 	}
// }
// marketCurrencyLastPrice[marketName] = last

// high, _ := strconv.ParseFloat(tick["high_24h"].(string), 64)
// low, _ := strconv.ParseFloat(tick["low_24h"].(string), 64)
// open, _ := strconv.ParseFloat(tick["open_24h"].(string), 64)
// volume, _ := strconv.ParseFloat(tick["volume_24h"].(string), 64)
// change := last - open
// changePercent := change / open

// currencySymbol := entity.MarketNameToCurrencySymbolMap[marketName]
// baseSymbol := entity.MarketNameToBaseNameMap[marketName]

// val := entity.QuoteTick{baseSymbol, currencySymbol, "gdax-" + currencySymbol + "-" + baseSymbol, open, high, low, last, volume, time.Now().In(loc), change, changePercent}
// entity.QuoteUpdates <- val
