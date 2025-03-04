package channels

import (
	"fmt"
)

func run(counts chan int) {
	for i := 1; i <= 50; i++ {
		counts <- i // Send value to the channel
	}
	close(counts) // Close the channel when done
}

func Channels() {
	counts := make(chan int)

	go run(counts) // Run the function in a goroutine

	fmt.Println(counts, "sdf")

	// Receive values from the channel
	for value := range counts {
		fmt.Println("Received:", value)
	}

}
