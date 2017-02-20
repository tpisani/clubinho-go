package main

import (
	"fmt"
	"time"
)

func expensive() {
	time.Sleep(1 * time.Second)
	fmt.Println("expensive")
}

func fast() {
	fmt.Println("fast")
}

func main() {
	go expensive()
	go fast()
	time.Sleep(2 * time.Second)
}
