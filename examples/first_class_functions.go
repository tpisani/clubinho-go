package main

import (
	"fmt"
)

func main() {
	hello := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}
	hello("John Doe")
}
