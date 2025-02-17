package mystruct

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {

	p := Person{name: name, age: 24}
	return &p

}

func Struct() {
	myPerson := newPerson("Harry")
	fmt.Println(*myPerson)

	fmt.Println(Person{name: "Ritesh", age: 21})

	fmt.Println(&Person{name: "Ann", age: 40})

	dog := struct {
		name string
		age  int
	}{
		name: "jonny",
		age:  2,
	}

	fmt.Println(dog)
}
