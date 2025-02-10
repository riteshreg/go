package functions

import "fmt"

func sum(nums ...int) int {
	fmt.Println("nums", nums)

	var sum int = 0

	for _, value := range nums {
		sum += value
	}

	return sum
}

func Variadic() {
	add := sum(4, 5)
	fmt.Println(add)

	nums := []int{1, 2, 3, 4}

	sum(nums...)

}
