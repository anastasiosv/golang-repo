package main

import (
	"fmt"
)

// send and receive chans
// concurrent code
func main() {
	c := make(chan int)

	//send
	go foo(c)

	// receive
	//if i had
	// go bar(c) then
	// bar func it would NOT
	// have the time to run
	// thus I am removing it
	bar(c)

	//
	fmt.Println("about to exit")

}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}

//cr := make(<-chan int) //receive
//cs := make(chan <- int) // send
