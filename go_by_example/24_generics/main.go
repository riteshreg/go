package generics

import (
	"fmt"
)

func PrintSlice[T int | string](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func Min[T int](s []T) T {
	if len(s) == 0 {
		return 0
	}

	min := s[0]

	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}

func Generic() {
	arr1 := []string{"ritesh", "hary", "shyam", "krishna"}
	PrintSlice(arr1)

	arr2 := []int{1, 2, 6, 43, 22, 11, 33, 55}
	// PrintSlice(arr2)

	outpp := Min(arr2)

	fmt.Println(outpp)
}
