package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		// Simulate slow operation
		time.Sleep(3 * time.Second)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
	}
}
