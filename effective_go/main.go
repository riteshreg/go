package main

import (
	"fmt"
)

// As an example, there's no need to spend time lining up the comments
// on the fields of a structure. Gofmt will do that for you. Given the declaration
type T struct {
	name string // this of the person
	age  int    //age of the person
}

func main() {
	fmt.Println("x")
}

func init() {
	fmt.Println("ran")
}
