package moretype

import "fmt"

func Pointer() {
	var x int = 4
	var ptx = &x

	fmt.Println(*ptx)
}
