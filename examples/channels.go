package main

import (
	"fmt"
)

func seq(n int) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func main() {
	s := seq(10)
	fmt.Println(<-s, <-s)
}
