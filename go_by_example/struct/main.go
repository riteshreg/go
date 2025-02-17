package mystruct

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name, age: 24}
	return &p

}

func Struct() {
	myPerson := newPerson("Harry")
	fmt.Println(*myPerson)

	fmt.Println(person{name: "Ritesh", age: 21})

	fmt.Println(&person{name: "Ann", age: 40})

	dog := struct {
		name string
		age  int
	}{
		name: "jonny",
		age:  2,
	}

	fmt.Println(dog)
}
