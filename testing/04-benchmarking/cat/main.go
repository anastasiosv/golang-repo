package main

import (
	"awesomeProject/go-fundamentals/testing/04-benchmarking/cat/mystr"
	"fmt"
	"strings"
)

const s ="We ask ourselves, who am i to be brilliant"

var xs []string

func main(){
	xs = strings.Split(s," ")
	for _,v := range xs{
		fmt.Println(v)

	}
	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))

}
