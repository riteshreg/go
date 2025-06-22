package main

import "fmt"

func zeroval(num int) {
	num = 0
}

func zeroptr(num *int) {
	*num = 0
}

func main() {
	v := 10
	fmt.Println("initial:", v)

	zeroval(v)
	fmt.Println("zeroval:", v)

	zeroptr(&v)
	fmt.Println("zeroptr:", v)

}
