package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			x := counter
			runtime.Gosched()
			x++
			counter = x

			fmt.Println("1 - counter", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end is: ", counter)
}
