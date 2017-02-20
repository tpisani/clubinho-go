package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("missing.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("bytes:", b)
}
