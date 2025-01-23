package moretype

import "fmt"

func SliceFromTheBlog() {

	x := append(make([]rune, 10), 'a', 'b')
	for _, r := range x {
		fmt.Printf("%c \n", r)
	}

	// b := [...]int{1, 2, 3, 4, 5}
	// b = append(b, 10) // can not do this in array

	// var s []int

	// s := make([]int, 4)

	// if s == nil {
	// 	fmt.Println("hi")
	// }

	// fmt.Println(s)

	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	fmt.Println(e) // a, d
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	fmt.Println(d)
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}
}
