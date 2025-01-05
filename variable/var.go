// Variables
// The var statement declares a list of variables; as in function argument lists, the type is last.

// A var statement can be at package or function level. We see both in this example.

package variable

import "fmt"

var c, nodejs, golang, java bool

func Var() {
	var value int
	fmt.Println(value, c, nodejs, golang, java) //the expected output is
	//0    false, false, false, false
}
