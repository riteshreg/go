package flowcontrol

import "fmt"

func Defer() {
	fmt.Println("first")
	defer fmt.Println("last")
	fmt.Println("middle")
}
