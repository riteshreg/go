package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r *rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	a1 := rect{width: 20, height: 10}
	fmt.Println(a1.perim())
	fmt.Println(a1.area())

}
