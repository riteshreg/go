package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	fmt.Println(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

}
