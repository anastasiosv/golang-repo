package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex
	// mutex is useful to prevent race conditions
	// locking and unlocking the resource
	// thus you will not be rejected
	// when you trie to do access it

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second * 2)
			runtime.Gosched() // doing the same
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("count: ", counter)
}
