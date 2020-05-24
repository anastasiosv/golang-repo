package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Start - GoRoutines:", runtime.NumGoroutine())

	var counter int64 // package atomic

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	//var mu sync.Mutex
	// mutex is useful to prevent race conditions
	// locking and unlocking the resource
	// thus you will not be rejected
	// when you trie to do access it

	for i := 0; i < gs; i++ {
		go func() {
			//using atomic to avoid
			// race conditions
			atomic.AddInt64(&counter, 1)
			runtime.Gosched() // doing the same
			fmt.Println("\tCounter:\t", atomic.LoadInt64(&counter))
			//mu.Lock()
			//time.Sleep(time.Second * 2)
			//mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("End - GoRoutines:", runtime.NumGoroutine())
	fmt.Println("count: ", counter)
}
