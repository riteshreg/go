// Functions continued
// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

// In this example, we shortened

// x int, y int
// to

// x, y int

package functions

import (
	"fmt"
)

// we can do better we can omit the type of a
func Sum(a int, b int) {
	var total int = a + b
	fmt.Println(total)
}

// ✅
func Subtract(a, b int) int {
	return a - b
}
