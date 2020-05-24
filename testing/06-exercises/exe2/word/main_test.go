package word

import (
	"awesomeProject/go-fundamentals/testing/06-exercises/exe2/quote"
	"fmt"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	xs := Count("My name is Bond James Bond")
	if xs != 6 {
		t.Error("Expected:\t",6,"\nGot:\t",xs)
	}
}

func TestUseCount(t *testing.T) {
	xs := UseCount("test")
	dictionary := map[string]int{"test":1}

	if reflect.DeepEqual(xs,dictionary)==false{
		t.Error("Eroorrrrrrr")
	}
}



func ExampleCount() {
	s := "My name is Bond James Bond"
	fmt.Println(Count(s))
	// Output:
	// 6
}
func ExampleUseCount() {
	fmt.Println(UseCount("My name is Bond James Bond"))
	// Output:
	// map[Bond:2 James:1 My:1 is:1 name:1]
}
func BenchmarkCount(b *testing.B) {
	for i:=0;i<b.N;i++{
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i:=0;i<b.N;i++{
	UseCount(quote.SunAlso)
	}
}

