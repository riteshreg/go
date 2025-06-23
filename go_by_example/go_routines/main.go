package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("go-routine")

	time.Sleep(time.Second)
	fmt.Println("done")
}
