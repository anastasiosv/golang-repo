package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10
	y := 10
	c := gen(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}
	fmt.Println("Routines: ", runtime.NumGoroutine())
}

func gen(x, y int) <-chan int {
	c := make(chan int)

	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
		fmt.Println("Routines: ", runtime.NumGoroutine())
	}
	return c
}

//func main(){
//	c := make(chan int)
//	go gen(c)
//	rec(c)
//	fmt.Println("bye")
//}
//
//func gen(c chan int){
//	var wg sync.WaitGroup
//
//	const goroutines = 10
//	wg.Add(goroutines)
//
//	for i:=0;i<goroutines;i++{
//		go func(){
//
//			for i:=0;i<10;i++{
//				c <- i
//			}
//			close(c)
//		}()
//
//	}
//	wg.Wait()
//
//	close(c)
//
//}
//
//func rec(c chan int){
//	for i:= range c{
//		fmt.Println(i)
//	}
//}
