package socketsclient

func WebsocketList() map[string]string {
	websocketList := map[string]string{
		"gdax": "wss://ws-feed.gdax.com",
	}

	return websocketList
}
