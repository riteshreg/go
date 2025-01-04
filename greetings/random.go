package greetings

import (
	"fmt"
	"math/rand"
)

func HelloRandomNumber() {
	random_num := rand.Intn(500)
	message := fmt.Sprintf("Hello world from %d", random_num)
	fmt.Println(message)
}
