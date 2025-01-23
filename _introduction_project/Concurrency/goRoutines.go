package concurrency

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func GoRoutine() {
	var wg sync.WaitGroup
	wg.Add(1) // Number of goroutines to wait for
	go printNumbers(&wg)
	wg.Wait() // Wait until all goroutines are done
	fmt.Println("Main function exiting...")
}
