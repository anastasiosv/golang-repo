package main

import (
	"fmt"
)

func main() {
	//	solution 2
	c := make(chan int, 1)

	c <- 42

	//solution 1
	//c := make(chan int)
	//go func(){
	//	c <- 42
	//}()

	fmt.Println(<-c)
}
