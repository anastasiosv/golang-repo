package dog

import (
	"fmt"
	"testing"
)
func TestYears(t *testing.T) {
	c := Years(10)
	if c != 70 {
		t.Error("Expected:\t",70,"\nGot:\t",c)
	}
}

func TestYearsTwo(t *testing.T) {
	c := YearsTwo(7)
	if c != 49 {
		t.Error("Expected:\t",49,"\nGot:\t",c)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i :=0; i<b.N; i++{
		Years(i)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i:=0;i<b.N;i++{
		YearsTwo(i)
	}
}