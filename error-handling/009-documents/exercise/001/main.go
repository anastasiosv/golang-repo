package main

import (
	"awesomeProject/go-fundamentals/error-handling/009-documents/exercise/001/dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	ruby := canine{
		name: "Ruby",
		age:  dog.Years(5),
	}

	fmt.Println(ruby)
}
