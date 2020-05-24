package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) // send
	//cs <- 42
	//cs <- 43
	//// you can't assign
	// since it's only int
	fmt.Println("S=====S")
	fmt.Printf("%T\n", c)
	fmt.Println("=====")
	fmt.Printf("%T\n", cr)
	fmt.Println("=====")
	fmt.Printf("%T\n", cs)

	// general to specific assigns
	cr = c
	cs = c

	fmt.Println("G=====S")
	fmt.Printf("%T\n", c)
	fmt.Println("=====")
	fmt.Printf("%T\n", (<-chan int)(c))
	fmt.Println("=====")
	fmt.Printf("%T\n", (chan<- int)(c))

}
