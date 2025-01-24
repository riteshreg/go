package constant

import "fmt"

const myname string = "Ritesh Khadka"

func Const() {
	fmt.Println(myname)

	const n = 500000000
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
}
