package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("That's all folks")

}

//package main
//
//import (
//	"fmt"
//)
//
//func main(){
//	c := gen()
//	receive(c)
//	fmt.Println("That's all folks")
//
//}
//
//func receive(c <-chan int){
//	for v :=range c{
//		fmt.Println(v)
//	}
//}
//
//
//func gen() <- chan int{
//	c := make(chan int)
//
//	go func() {
//		for i := 0; i < 100; i++ {
//			c <- i
//		}
//		close(c)
//	}()
//	return c
//}
