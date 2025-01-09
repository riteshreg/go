package flowcontrol

import (
	"fmt"
	"math"
)

func Pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println(v)
	}
	return lim
}

func Sqart(x float64) float64 {
	z := 1.0

	for i := 0; i < 100; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println(i, z)
	}

	fmt.Println(z)
	return z
}
