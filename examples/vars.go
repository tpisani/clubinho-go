package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	var x, y int
	x = 5
	y = 10
	fmt.Println(x, "+", y, "equals", add(x, y))
}
