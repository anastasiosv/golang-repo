package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x before: ", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)

}

//everything in go is pass by value

func foo(y *int) {
	fmt.Println("y before: ", y)
	fmt.Println("y before: ", *y)
	*y = 43
	fmt.Println("y after: ", y)
	fmt.Println("y after: ", *y)

}

//
////v1
//func main(){
//	x := 0
//	foo(x)
//	fmt.Println("hoooola", x)
//}
////everything in go is pass by value
//
//
//func foo(y int){
//	fmt.Println(y)
//	y = 43
//	fmt.Println(y)
//}

// vol 2
//func main(){
//	x := 0
//	fmt.Println("x before: ",&x)
//	fmt.Println("x before",x)
//	foo(&x)
//	fmt.Println("x after", &x)
//	fmt.Println("x after", x)
//
//}
////everything in go is pass by value
//
//
//func foo(y *int){
//	fmt.Println("y before: ",y)
//	fmt.Println("y before: ",*y)
//	*y = 43
//	fmt.Println("y after: ",y)
//	fmt.Println("y after: ",*y)
//
//}
