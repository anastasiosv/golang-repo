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
	buckets := make([][]string, 12)
	// create slices to hold words
	for i:=0; i<12 ; i++{
		buckets = append(buckets, []string{})
	}
	// loop over the words
	for sc.Scan(){
		word := sc.Text()
		n := HashBucket(word,12)
		buckets[n] = append(buckets[n], word)

	}
	for i:=0;i<12;i++ {
		fmt.Println(i," - ",len(buckets[i]))
	}
	// Print the words of one bucket
	//fmt.Println(buckets[3])
}


func HashBucket(word string, buckets int) int {
	var sum int
	for _,v := range word {
		sum += int(v)
	}
	return sum % buckets
}