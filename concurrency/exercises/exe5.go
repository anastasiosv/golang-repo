package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	var counter int64
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("1 - counter", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end is: ", counter)
}
