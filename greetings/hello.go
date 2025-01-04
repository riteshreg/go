package greetings // ✅ Declare the package name

import (
	"fmt"
	"math/rand"
)

// Hello function (Exported because it starts with a capital letter)
func Hello(name string) {
	var message string = fmt.Sprintf("Hey %s, Happy New Year!!! 👌👌", name)
	fmt.Println(message)
}

func HelloRandomNumber() {
	random_num := rand.Intn(500)
	message := fmt.Sprintf("Hello world from %d", random_num)
	fmt.Println(message)
}
