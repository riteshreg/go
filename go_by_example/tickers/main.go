package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Microsecond * 500)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	ticker.Stop()

	done <- true

	fmt.Println("Ticker stopped")

}
