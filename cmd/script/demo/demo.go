package main

import (
	"fmt"
	"time"
)

type Req any

func handle(r Req) { fmt.Println(r.(int)) }

const RateLimitPeriod = time.Minute
const RateLimit = 200 // 一分钟最多 200

func handleReqs(reqs <-chan Req) {
	quotas := make(chan time.Time, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
				fmt.Println("send")
			default:
				fmt.Println("drop")
			}
		}
	}()

	for r := range reqs {
		<-quotas
		go handle(r)
	}
}

func main() {
	requests := make(chan Req)
	go handleReqs(requests)

	time.Sleep(time.Minute)
	for i := 0; ; i++ {
		requests <- i
	}
}
