package main

import (
	"fmt"
)

// send and receive chans
// concurrent code
func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)
	// receive
	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func receive(e, o, q <-chan int) {
	for { //infinate loop
		select { //like switch
		case v := <-e:
			fmt.Println("from even ch:", v)

		case v := <-o:
			fmt.Println("from odd ch:", v)

		case v := <-q:
			fmt.Println("from quit ch", v)
			//close(q)
			return
		}
	}
}

// send
func send(e, o, q chan<- int) {

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}

	}
	// if i close them
	// some of them at the end have 0s
	//close(e)
	//close(o)
	q <- 0
}
