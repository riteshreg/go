package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name, age: 21}
	return &p
}

func main() {
	p := *newPerson("Ritesh Khadka")
	fmt.Println(p.age)

}
