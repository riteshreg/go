// For is Go's "while"
package flowcontrol

import "fmt"

func ContinuedLoop() {
	sum := 1
	for sum < 5024 {
		sum = sum + sum
	}

	fmt.Println(sum)
}
