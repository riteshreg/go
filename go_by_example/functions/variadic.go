package main

import "fmt"

func Sum(nums ...int) int {
	var total = 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	var sum = Sum(10, 20, 30)
	fmt.Println(sum)
}
