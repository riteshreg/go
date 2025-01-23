package moretype

import "fmt"

func SliceLiterals() {
	var arr = [2]bool{false, true} // so this is the array literals we are declearing its size
	fmt.Println(arr)
	//but in slice literals we do not have to specify a size of array
	var arr2 = []bool{true, false, true, false}
	fmt.Println(arr2)

	t := []struct {
		i   int
		yes bool
	}{
		{1, false},
		{i: 2, yes: true},
		{yes: false, i: 5},
	}

	fmt.Println(t)
}
