package main

import (
	"fmt"
)

// send and receive chans
// concurrent code
func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		// important step to close
		// the channel
	}()
	// receive

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")

}

// send
func foo(c chan<- int) {

	close(c)
}
