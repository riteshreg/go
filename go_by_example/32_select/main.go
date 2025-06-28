package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "sending from c1"
	}()

	go func() {
		time.Sleep(time.Second * 4)
		c2 <- "sending from c2"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
