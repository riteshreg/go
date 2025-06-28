package main

import "fmt"

// Example 1: Basic embedding with multiple embedded structs
type Animal struct {
	Name string
	Age  int
}

func (a Animal) MakeSound() string {
	return "Some sound"
}

type Dog struct {
	Animal        // Embed Animal
	Breed  string // Dog-specific field
}

func (d Dog) MakeSound() string {
	return "Woof!"
}

// Example 2: Embedding with method overriding
type Vehicle struct {
	Brand string
	Model string
}

func (v Vehicle) GetInfo() string {
	return fmt.Sprintf("%s %s", v.Brand, v.Model)
}

type Car struct {
	Vehicle        // Embed Vehicle
	Color   string // Car-specific field
}

func (c Car) GetInfo() string {
	return fmt.Sprintf("%s in %s color", c.Vehicle.GetInfo(), c.Color)
}

// Example 3: Embedding interfaces
type Speaker interface {
	Speak() string
}

type Human struct {
	Name string
}

func (h Human) Speak() string {
	return fmt.Sprintf("Hello, I'm %s", h.Name)
}

type Robot struct {
	Human        // Embed Human (which implements Speaker)
	Model string // Robot-specific field
}

// Example 4: Embedding with field name conflicts
type Point struct {
	X, Y int
}

type Circle struct {
	Point  // Embed Point
	Radius int
}

type Rectangle struct {
	Point  // Embed Point
	Width  int
	Height int
}

func main() {
	fmt.Print("=== Struct Embedding Examples ===\n")

	// Example 1: Basic embedding
	fmt.Println("1. Basic Embedding:")
	dog := Dog{
		Animal: Animal{Name: "Buddy", Age: 3},
		Breed:  "Golden Retriever",
	}
	fmt.Printf("Dog: %s, Age: %d, Breed: %s\n", dog.Name, dog.Age, dog.Breed)
	fmt.Printf("Sound: %s\n\n", dog.MakeSound())

	// Example 2: Method overriding
	fmt.Println("2. Method Overriding:")
	car := Car{
		Vehicle: Vehicle{Brand: "Toyota", Model: "Camry"},
		Color:   "Blue",
	}
	fmt.Printf("Car Info: %s\n\n", car.GetInfo())

	// Example 3: Interface satisfaction through embedding
	fmt.Println("3. Interface Satisfaction:")
	robot := Robot{
		Human: Human{Name: "R2D2"},
		Model: "Protocol Droid",
	}
	fmt.Printf("Robot speaks: %s\n", robot.Speak())
	fmt.Printf("Robot model: %s\n\n", robot.Model)

	// Example 4: Field access
	fmt.Println("4. Field Access:")
	circle := Circle{
		Point:  Point{X: 10, Y: 20},
		Radius: 5,
	}
	fmt.Printf("Circle at (%d, %d) with radius %d\n", circle.X, circle.Y, circle.Radius)
	fmt.Printf("Same as: (%d, %d)\n\n", circle.Point.X, circle.Point.Y)

	// Example 5: Anonymous embedding
	fmt.Println("5. Anonymous Embedding:")
	type Anonymous struct {
		int    // Embed primitive type
		string // Embed primitive type
	}

	anon := Anonymous{42, "hello"}
	fmt.Printf("Anonymous: %d, %s\n", anon.int, anon.string)
}
