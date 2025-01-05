package variable

import (
	"fmt"
)

var i, j int = 1, 2

func WithInit() {
	var c, nodejs, x, y = false, false, 1, 2
	fmt.Println(c, nodejs, x, y, i, j)
}
