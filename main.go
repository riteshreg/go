package main

import (
	"fmt"
	"golang/functions"
	"golang/greetings"
)

func main() {
	greetings.Hello("Ritesh Khadka")
	greetings.HelloRandomNumber()

	// function package <<
	functions.Sum(1, 2)
	var subtract int = functions.Subtract(3, 2)
	fmt.Println(subtract)

	a, b := functions.Swap("Ritesh", "Khadka")
	fmt.Printf("a=%s and b=%s \n", a, b)

	first, secod := functions.Split(10)
	fmt.Printf("first=%d and second=%d", first, secod)

	// function package >>

}
