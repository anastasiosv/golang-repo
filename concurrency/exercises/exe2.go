package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) Speak() {
	fmt.Println("My name is ", p.name)
}

type human interface {
	Speak()
}

func saySomething(h human) {
	h.Speak()
}

func main() {
	fmt.Println("hey")
	p := person{
		name: "James Bond",
		age:  42,
	}

	saySomething(&p)

}
