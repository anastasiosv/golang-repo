package main

import (
	"fmt"
)

type person struct {
	age int
}

func changeMe(p *person) {
	(*p).age += 1 //dereferencing
	// p.first
}

func main() {
	p := person{42}
	fmt.Println(p.age)
	changeMe(&p)
	fmt.Println(p.age)

}
