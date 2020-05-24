package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T){
	s := Greet("James")
	if s != "Welcome James,my old friend"{
		t.Error("got", s, "expected ","Welcome James,my old friend")
	}

}

func ExampleGreet(){
	fmt.Println(Greet("James"))
	// Output:
	// Welcome James,my old friend
}

func BenchmarkGreet(b *testing.B){
	for i:= 0; i< b.N; i++{
		Greet("James")
	}
}