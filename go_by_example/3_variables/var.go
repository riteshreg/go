package variables

import "fmt"

func Var() {
	var a string = "Hello, World!"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	const d = true
	fmt.Println(d)

	e := "short"
	fmt.Println(e)
}
