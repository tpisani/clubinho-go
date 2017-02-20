package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	x := 5
	y := 10
	result := add(x, y)
	fmt.Println(x, "+", y, "equals", result)
}
