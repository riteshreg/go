package moretype

import "fmt"

func SliceBound() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)

}
