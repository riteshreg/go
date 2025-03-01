package goroutines

import (
	"fmt"
	"time"
)

// func f(from string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(from, ":", i)
// 	}
// }

func display(str string) {
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
	}
}

func GoRoutine() {
	// f("direct")

	// go f("goroutine")

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// time.Sleep(time.Second)
	// fmt.Println("done")

	go display("Hello, Goroutine!") // Runs concurrently

	display("Hello, Main!")

	// time.Sleep(500 * time.Millisecond)

	go func(str string) {
		for i := 0; i < 3; i++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(str)
		}
	}("Hi My Name is Ritesh and I am speaking from go Anonymous  function")

}
