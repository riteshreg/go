package array

import "fmt"

func Array() {
	a := [3]int{5, 5, 4}

	for i, v := range a {
		fmt.Println(i, v)
	}

}
