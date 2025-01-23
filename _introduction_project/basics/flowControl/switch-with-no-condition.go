package flowcontrol

import (
	"fmt"
	"time"
)

func SwitchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Print("Good Afternoon")
	default:
		fmt.Print("Good Evening")
	}
}
