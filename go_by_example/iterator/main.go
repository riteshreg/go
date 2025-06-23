package main

import (
	"fmt"
	"iter"
)

// Fibonacci iterator
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	// Range over the Fibonacci iterator
	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
