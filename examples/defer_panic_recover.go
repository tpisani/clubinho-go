package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("something went wrong:", err)
		}
	}()
	panic("I'm panicking")
}
