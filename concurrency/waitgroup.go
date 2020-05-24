package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()

	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1) // wait for 1 thing
	//go foo() // run a routine
	go foo()
	bar()
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("time of exec:", elapsed)
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
