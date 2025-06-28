package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	fmt.Println("waiting")
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// stop2 := timer2.Stop()
	// if stop2 {
	// 	fmt.Println("Timer 2 stopped")
	// }

	<-time.NewTimer(2 * time.Second).C

}
