// go only has one loop which is obviously  for loop

package flowcontrol

import "fmt"

func SumWithLoop() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum++
	}

	fmt.Println(sum)
}
