package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)  // value
	fmt.Println(&a) // address

	fmt.Printf("%T\n", a)  // int
	fmt.Printf("%T\n", &a) // *int
	// two very seperate types

	// you can't do -> var b int = &a
	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(&b)

	c := &a
	fmt.Println("c: \n", c) //address
	fmt.Println(*c)         //42
	fmt.Println(*&a)        //42
	//* -> value stored on an address
	// & -> address

	*b = 43 //share the same address
	fmt.Println(a)
}
