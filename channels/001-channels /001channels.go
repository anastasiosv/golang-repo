package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	// channels block
	// here we need to fix the c
	// we need a go routine to make it run

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
