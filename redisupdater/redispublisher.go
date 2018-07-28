package redispublisher

import (
	"fmt"
)

func PublishToRedis(ch chan interface{}) {
	fmt.Println("PUBLISHING TO REDIS")
}
