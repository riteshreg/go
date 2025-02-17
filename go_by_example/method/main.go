package method

import (
	"fmt"
)

// Defining a struct
type person struct {
	name string
	age  int
}

// Defining a method with struct receiver
func (p person) dispaly() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}

type number int

func (n number) sqart() number {
	return n * n
}

func (p *person) changeName(newName string) {
	p.name = newName
}

func Methods() {

	mynewperson := person{name: "ritesh", age: 21}

	mynewperson.dispaly()

	mynewperson.changeName("harry")

	(mynewperson).dispaly()

	var number number = 10
	fmt.Println(number.sqart())

}
