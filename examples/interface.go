package main

import (
	"fmt"
)

type Speaker interface { // HLINTERFACE
	Speak() string // HLINTERFACE
} // HLINTERFACE

type Developer struct{}

func (d Developer) Speak() string { // HLINTERFACE
	return "There are only 10 types of people..." // HLINTERFACE
} // HLINTERFACE

func saySomething(s Speaker) { // HLINTERFACE
	fmt.Println(s.Speak()) // HLINTERFACE
} // HLINTERFACE

func main() {
	saySomething(Developer{})
}
