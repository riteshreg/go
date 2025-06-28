package main

import "fmt"

func recursive(n int) int {
	if n == 0 {
		return 0
	}

	return recursive(n - 1)
}

func main() {
	fmt.Println(recursive(7))
}
