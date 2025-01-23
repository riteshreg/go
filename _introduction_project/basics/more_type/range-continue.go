package moretype

import "fmt"

func RangeContinue() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = i*2 + 1
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
