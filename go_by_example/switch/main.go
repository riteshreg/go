package switch_statement

import (
	"fmt"
	"time"
)

func Switch() {
	i := 3

	fmt.Printf("i is %d \n", i)
	switch i {
	case 1:
		fmt.Println("I is 1")
	case 2:
		fmt.Println("I is 2")
	case 3:
		fmt.Println("I is 3")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now().Hour()

	switch {
	case t < 12:
		fmt.Println("Its moring")
	default:
		fmt.Println("Its evening")
	}

}
