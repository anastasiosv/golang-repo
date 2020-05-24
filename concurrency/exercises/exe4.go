package main

import (
	"fmt"
	"runtime"
	"sync"
)

//var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	counter := 0
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			//mu.Lock()
			x := counter
			runtime.Gosched()
			x++
			counter = x
			//if my unlock is here, i do have a datarace
			// because of reading
			fmt.Println("1 - counter", counter)
			//mu.Unlock()
			//unlock should be before the done message
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end is: ", counter)
}
