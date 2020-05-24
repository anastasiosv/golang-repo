package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Main")
	wg.Add(2)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("aaa: ", i)
		}
		wg.Done()

	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("bbb: ", i)
		}
		wg.Done()
	}()
	fmt.Println("middle CPUs\t", runtime.NumCPU())
	fmt.Println("middle Goroutines\t", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

}
