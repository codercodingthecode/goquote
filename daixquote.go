package main

import (
	"fmt"

	"github.com/codercodingthecode/goquote/exchanges"
)

func main() {

	chStatus := make(chan string)

	go exchanges.SubscribeToGDAX(chStatus)

	fmt.Println("THERE WAS AN ERROR ---> ", <-chStatus)

}
