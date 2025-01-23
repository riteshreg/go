package moretype

import "fmt"

func Slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	s := primes[1:5] // include first and exclude last which mean 1:5 means [3, 5, 7, 11]
	fmt.Println(s)
}
