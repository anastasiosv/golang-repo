package main

import (
	"fmt"
)

// send and receive chans
// concurrent code
func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(even, odd, quit)
	// receive
	receive(even, odd, quit)

	fmt.Println("about to exit")
}

// send
func send(even, odd chan<- int, quit chan<- bool) {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)

}

func receive(even, odd <-chan int, quit <-chan bool) {
	for { //infinate loop
		select { //like switch
		case v := <-even:
			fmt.Println("from even ch:", v)
		case v := <-odd:
			fmt.Println("from odd ch:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok", i, ok)
			} else {
				fmt.Println("from comma ok", i)
			}
			return
		}
	}
}
