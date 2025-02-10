package slice

import "fmt"

func Slice() {
	// var s []string

	// fmt.Println(s, s == nil, len(s) == 0)

	// s = make([]string, 3)
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// s[3] = "d"
	// s[4] = "e"

	// fmt.Println("set:", s)
	// fmt.Println("get:", s[2])

	nums := make([]int, 3, 5)           // length: 3, capacity: 5
	fmt.Println("Slice:", nums)         // [0 0 0]
	fmt.Println("Length:", len(nums))   // 3
	fmt.Println("Capacity:", cap(nums)) // 5

	nums = append(nums, 4) // Append within capacity
	fmt.Println("After Append:", nums)
	fmt.Println("New Length:", len(nums))   // 4
	fmt.Println("New Capacity:", cap(nums)) // 5

	nums = append(nums, 5, 6) // Exceeds capacity, new allocation happens
	fmt.Println("After More Appends:", nums)
	fmt.Println("New Length:", len(nums))   // 6
	fmt.Println("New Capacity:", cap(nums)) // Increased automatically

}
