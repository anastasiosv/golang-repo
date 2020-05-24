package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.mieliestronk.com/corncob_lowercase.txt")
	if err != nil {
		log.Fatal(err)
	}
	//scan the page
	sc := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// split function for scanning op
	sc.Split(bufio.ScanWords)
	// create slice to hold counts
	buckets := make([]int, 12)
	// loop over the words
	for sc.Scan(){
		n := HashBucket(sc.Text(),12)
		buckets[n]++
	}
	fmt.Println(buckets)

}


func HashBucket(word string, buckets int) int {
	var sum int
	for _,v := range word {
		sum += int(v)
	}
	return sum % buckets
}