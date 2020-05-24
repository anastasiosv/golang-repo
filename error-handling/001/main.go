//test my test
package main

import (
	"fmt"
)

//it's a simple hello world with error handling
func main() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(n)
}
