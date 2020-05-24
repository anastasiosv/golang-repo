package main

import (
	"awesomeProject/go-fundamentals/testing/03-examples-tests/acdc"
	"fmt"
)

func main() {
	x := acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(x)
}
