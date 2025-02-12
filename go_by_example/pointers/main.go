package pointers

import "fmt"

func xx(y *int) {
	*y = 10
	fmt.Println("inside", *y)
}

func myarr(arr_ptr *[3]int) {
	fmt.Println(*arr_ptr)
}

func Pointers() {
	var x int = 10
	var ptr1 *int = &x

	fmt.Println(x)     // output: 10
	fmt.Println(ptr1)  // output: 0xc0000140a8
	fmt.Println(*ptr1) // output: 10

	var ptr *int = new(int)

	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = 10

	fmt.Println(*ptr)

	ptr = nil // nill will kill the memory address

	y := 2
	fmt.Println("outside before", y)
	xx(&y)
	fmt.Println("outside after", y)

	arr := [3]int{1, 2, 3}
	ptr2 := &arr[0]
	fmt.Println(*ptr2) // output: 1
	*ptr2++
	fmt.Println(*ptr2) // output: 2

	arr2 := [3]int{1, 2, 3}

	myarr(&arr2)

}
