package flowcontrol

import (
	"fmt"
	"time"
)

func SwitchEvaluation() {
	// as of: today is Thursday
	today := time.Now().Weekday()
	// we are getting Thursday at today variable

	switch time.Saturday { // we expect to be saturday
	case today + 0: //thursday + 0 is thursday
		fmt.Println("Today")
	case today + 1: //thursday + 1 is friday
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("in 2 days")
	default:
		fmt.Println("to far")
	}

}
