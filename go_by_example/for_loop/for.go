package forloop

import "fmt"

func For() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println("Hello, World!", i)
	// }

	// i := 0

	// for i <= 7 {
	// 	fmt.Println("Hello S", i)
	// 	i++
	// }

	for i := range 6 {
		fmt.Println(i)
	}
}
