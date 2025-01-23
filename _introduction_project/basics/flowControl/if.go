package flowcontrol

import (
	"fmt"
	"math"
)

func Sqrt(num float64) string {
	if num < 0 {
		return Sqrt(-num) + "i"
	}

	return fmt.Sprint(math.Sqrt(num))
}
