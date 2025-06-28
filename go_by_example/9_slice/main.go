package main

import "fmt"

func main() {
	var s []string
	if s == nil {
		fmt.Println("s is nil")
	}

	d := make([]string, 0, 10)
	if d == nil {
		fmt.Println("d is nil")
	}
	fmt.Println(len(d))
}
