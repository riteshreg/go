package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "रितेश"
	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\n", s[i])
	}

	fmt.Println("rune count:", utf8.RuneCountInString(s))
}
