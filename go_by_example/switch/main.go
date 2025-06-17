package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	switch i {
	case 1:
		fmt.Println("I is 1")
	case 2:
		fmt.Println("I is 2")
	case 3:
		fmt.Println("I is 3")
	default:
		fmt.Printf("default: I is %d\n", i)
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

}
