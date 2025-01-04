package greetings // ✅ Declare the package name

import (
	"fmt"
)

// Hello function (Exported because it starts with a capital letter)
func Hello(name string) {
	var message string = fmt.Sprintf("Hey %s, Happy New Year!!! 👌👌", name)
	fmt.Println(message)
}
