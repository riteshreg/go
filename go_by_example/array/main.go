package main

import "fmt"

func main() {
	var names [2]string = [2]string{"Ritesh Khadka", "Nayan Regmi"}
	fmt.Println(names)
	fmt.Println("len", len(names))

	for _, name := range names {
		fmt.Println(name)
	}

	cities := [...]string{"Damak", "Birthmode", "Kathmandu"}

	for idx, city := range cities {
		fmt.Printf("%d is the index of %s\n", idx, city)
	}

}
