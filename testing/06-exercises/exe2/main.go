package main

import (
	"awesomeProject/go-fundamentals/testing/06-exercises/exe2/quote"
	"awesomeProject/go-fundamentals/testing/06-exercises/exe2/word"
	"fmt"

)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	m := word.UseCount(quote.SunAlso)
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(v, k)
	}
}