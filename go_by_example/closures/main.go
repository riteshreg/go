package closures

import "fmt"

func sum(val1 int) func(val2 int) int {
	return func(val2 int) int {
		return val1 + val2
	}
}

func Closures() {
	f1 := sum(10)
	fmt.Println(f1(11))
}
