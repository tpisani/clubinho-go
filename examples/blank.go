package main

import (
	"fmt"
)

func main() {
	names := []string{"John Doe", "Jane Doe"}
	for _, name := range names {
		fmt.Println(name)
	}
}
