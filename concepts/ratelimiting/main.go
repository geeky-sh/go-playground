package main

import (
	"fmt"
	"time"
)

/*
What I want to build:
1. Simulate the ratelimiting for 1 requests / 1 second
2. Simulation the bursty limiter of allowing 90 requests but then reduce it it 1 request / second
*/
func printReq(req int) {
	fmt.Printf("Request %d received at %s\n", req, time.Now())
}

func simulateReqs(cnt int) chan int {
	reqs := make(chan int, cnt)
	for i := 1; i <= cnt; i++ {
		reqs <- i
	}
	close(reqs)
	return reqs
}

func main() {
	reqs := simulateReqs(5)

	fmt.Println(" Rate Limiter")

	rateLimiter := time.Tick(1 * time.Second)

	for req := range reqs {
		<-rateLimiter
		printReq(req)
	}

	fmt.Println("Bursty Rate Limiter")

	freeReq := 90
	burstyRateLimiter := make(chan int, freeReq)

	for i := 1; i <= freeReq; i++ {
		burstyRateLimiter <- i
	}

	go func() {
		for t := range time.Tick(1 * time.Second) {
			burstyRateLimiter <- int(t.Unix())
		}
	}()

	reqs = simulateReqs(freeReq + 5)
	for req := range reqs {
		<-burstyRateLimiter
		printReq(req)
	}

}
