package main

import (
	"fmt"
	"time"
)

type customErr struct {
	msg  string
	when time.Time
}

func foo(myerr customErr) string {
	return fmt.Sprintf("myerror : %v at : %v", myerr.msg, myerr.when)
}

func main() {
	c := customErr{"oopsy", time.Now()}
	q := foo(c)
	fmt.Println(q)
}
