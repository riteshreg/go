package my_interface

import (
	"fmt"
)

// // Define the interface
// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// // Circle type that implements the Shape interface
// type Circle struct {
// 	radius float64
// }

// // Rectangle type that implements the Shape interface
// type Rectangle struct {
// 	length, width float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func (c Circle) Perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func (r Rectangle) Area() float64 {
// 	return r.length * r.width
// }

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.length + r.width)
// }

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
	age  int
}

func (d Dog) Speak() {
	fmt.Println("bhau bhau")
}

type Cat struct {
	name string
	age  int
}

func (c Cat) Speak() {
	fmt.Println("meow meow")
}

func MakeSound(s Speaker) {
	s.Speak()
}

func Interface() {
	var dog_one = Dog{name: "Tommy", age: 5}
	MakeSound(dog_one)

	var cat_one = Cat{name: "puku", age: 2}
	MakeSound(cat_one)
}
