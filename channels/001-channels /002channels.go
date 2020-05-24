package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 2)
	// this is successful because
	// i set a buffer number on the channel
	// equal to 1

	c <- 42
	c <- 43
	fmt.Println(<-c) //42
	fmt.Println(<-c) //43

}
