package moretype

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Range() {

	for i, v := range pow {
		// i is index
		//v is copy of the current value
		fmt.Println(i, v)
	}

}
