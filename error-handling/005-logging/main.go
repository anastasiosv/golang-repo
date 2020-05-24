package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		log.Fatal(err)
		panic(err)
	}
	//	log.Println("hello world")
}
