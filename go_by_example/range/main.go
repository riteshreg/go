package ranges

import "fmt"

func Range() {
	nums := []int{4, 5, 6}

	var sum int

	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	kvs := map[string]string{"a": "apple", "b": "ball", "c": "cat", "d": "dog"}

	for key, value := range kvs {
		fmt.Printf("key is %s and value is %s \n", key, value)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "abcd" {
		fmt.Println(i, c)
	}

}
