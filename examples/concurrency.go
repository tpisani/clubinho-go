package main

import (
	"fmt"
	"time"
)

func expensive() {
	time.Sleep(1 * time.Second)
	fmt.Println("expensive")
}

func cheap() {
	fmt.Println("cheap")
}

func main() {
	go expensive()
	go cheap()
	time.Sleep(2 * time.Second)
}
