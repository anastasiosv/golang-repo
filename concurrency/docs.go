package main

import (
	"fmt"
	"time"
)

func dosmth(x int) int {
	return x * 5
}

func main() {
	start := time.Now()
	ch := make(chan int)
	go func() {
		ch <- dosmth(5)
	}()
	fmt.Println(<-ch)
	elapsed := time.Since(start)
	fmt.Println("time of exec:", elapsed)
}
