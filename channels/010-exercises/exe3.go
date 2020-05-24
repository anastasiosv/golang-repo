package main

import (
	"fmt"
)

func main() {

	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		//i need a go func
		// because i can't close the channel
		// and return it after that
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
