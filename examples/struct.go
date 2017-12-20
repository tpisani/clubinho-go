package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) Speak() string { // HLMETHOD
	return fmt.Sprintf("My name is %s.", p.Name) // HLMETHOD
} // HLMETHOD

func main() {
	john := Person{Name: "John"}
	fmt.Println(john.Speak()) // HLMETHOD
}
