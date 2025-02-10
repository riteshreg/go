package maps

import (
	"fmt"
	"maps"
)

func Maps() {

	m := make(map[string]int, 4)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6

	fmt.Println(m)

	delete(m, "one")

	fmt.Printf("after deletion %v \n", m)

	clear(m)
	fmt.Println("map:", m)

	val, prs := m["two"]
	fmt.Println("prs:", prs, val)

	n1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n1)

	n2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(n1, n2) {
		fmt.Println("n == n2")
	}

}
