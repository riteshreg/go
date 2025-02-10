package multireturnvalue

import "fmt"

func vals() (int, int) {
	return 1, 2
}

func Multireturnvalue() {
	a, b := vals()
	fmt.Println(a, b)

	c, _ := vals()
	fmt.Println(c)

	_, d := vals()
	fmt.Println(d)

}
