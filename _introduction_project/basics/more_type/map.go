// Syntax
// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

package moretype

import "fmt"

func Map() {
	var a = map[string]string{"name": "ritesh", "age": "21"}
	var b = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	c := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
	fmt.Printf("c\t%v\n", c)

}

func MapWithMake() {

	// create map using make
	// var a = make(map[KeyType]ValueType)
	// b := make(map[KeyType]ValueType)
	var ax = make(map[string]string)
	ax["name"] = "Ritesh"
	fmt.Printf("ax\t%v\n", ax)

	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	// a is no longer empty
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	fmt.Println(a["brand"])
	val1, ok1 := a["brand"] // Checking for existing key and its value
	fmt.Println(val1, ok1)

	delete(a, "brand")
	fmt.Println(a)

	val2, ok2 := a["brand"] // Checking for existing key and its value
	fmt.Println(val2, ok2)

}
