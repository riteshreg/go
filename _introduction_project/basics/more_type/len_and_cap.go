package moretype

import "fmt"

func LenAndCap() {
	var a = [3]int{1, 2, 3}
	var b = a[0:2]
	o := fmt.Sprintf("a %v and b %v", a, b)
	fmt.Println(o)

	fmt.Printf("len %v \n", len(b))
	fmt.Printf("cap %v \n", cap(b))
}
