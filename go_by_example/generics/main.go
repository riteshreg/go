package generics

import "fmt"

func PrintSlice[T int | string](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func Generic() {
	arr1 := []string{"ritesh", "hary", "shyam", "krishna"}
	PrintSlice(arr1)

	arr2 := []int{1, 2, 6, 43, 22, 11, 33, 55}
	PrintSlice(arr2)
}
