package main

import "fmt"

func main(){
	n := HashBucket("Go",12)
	fmt.Println(n)
}

func HashBucket(word string, buckets int) int{
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}

////rune version
//func HashBucket(word string, buckets int32) int32{
//	letter := rune(word[0])
//	bucket := letter % buckets
//	return bucket
//}