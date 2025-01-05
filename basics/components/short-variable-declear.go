// Short variable declarations
// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

package components

import "fmt"

func ShortVariableDec() {
	var a, b int = 1, 2
	name := "Ritesh Khadka" //this is not allowed outside a function
	age := 21
	fmt.Println(a, b, name, age)
}
