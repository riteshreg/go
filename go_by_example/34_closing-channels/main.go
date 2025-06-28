package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 39
	ch <- 49

	close(ch)

	// ch <- 50 panic

	for val := range ch {
		fmt.Println(val)
	}
}
