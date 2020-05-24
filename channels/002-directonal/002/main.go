package main

import (
	"fmt"
)

func main() {

	c := make(chan<- int, 2)
	// you can only sent values

	c <- 42
	c <- 43
	// you can't pull out from chan
	// since it's only int
	fmt.Println(<-c) //42
	fmt.Println(<-c) //43

	fmt.Println("=====")
	fmt.Printf("%T\n", c)

}
