package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("Returned normally from F.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("hey G")
	g(0)
	fmt.Println("Returned normal gz")

}

func g(i int) {
	if i > 3 {
		fmt.Println("Panik")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
