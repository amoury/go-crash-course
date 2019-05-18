package main

import "fmt"

// Define Person
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Value receiver
func (p Person) greet() string {
	return "Hello , my name is " + p.firstName
}

// Pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	x := "Ansar"
	person1 := Person{x, "Memon", "Blore", "Male", 27}

	fmt.Println(person1.greet())

	person1.hasBirthday()

	fmt.Println(person1.age)
}
