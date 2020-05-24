package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 9, errors.New("nope: square root of negative num")
	}
	return 42, nil

}
