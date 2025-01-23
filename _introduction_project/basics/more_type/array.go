package moretype

import "fmt"

func Array() {
	var a [4]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	// init
	var b [2]string = [2]string{"Ram", "Shyam"}
	fmt.Println(b)

	var c = [3]bool{false, false, true}
	fmt.Println(c)

	d := [3]float32{3.3, 4.3, 6.6}
	fmt.Println(d)
}
