package main

import "fmt"

// design interfaces in Go
// instead of designing in terms of what kind of data our types can hold
// we design our abstractions in terms of what actions our types can execute

// an interface value is constructed of two words of data
// on word used to point to the actual data being held by that value
// another points to a method table for the value's underlying type

// a value of interface type can hold any value that implements those methods

// interface
type Animal interface {
	Speak() string
	// output type of Speak is a string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!" // receiver is of type Dog, the method imexplicitly implement interface Animal
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design Patterns!"
}

func main() {
	fmt.Println("Hello Gophers!")
	animals := []Animal{Dog{}, Cat{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
