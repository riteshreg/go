package moretype

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

func Struct() {
	// x := Vertex{x: 1, y: 2}
	// fmt.Println(x)

	v := Vertex{x: 1, y: 2}
	v.x = 1e4
	fmt.Println(v)
}
