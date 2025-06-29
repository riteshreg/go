package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}

	// chosing the requests chan now all the write on the chan is invalid
	// but we read from this closed chan
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for range 3 {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("done")
}
