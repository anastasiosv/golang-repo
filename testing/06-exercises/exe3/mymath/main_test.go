package mymath

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/007-documentation/01/mymath"
	"testing"
)


func TestCenteredAvg(t *testing.T) {
	type test struct {
		data []int
		answer float64
	}

	tests := []test{
		test{[]int{10,20,40,60,80},40,},
		test{[]int{2,4,6,8,10,12},7,},
		test{[]int{1,2,3,4,5,6,7,8,9}, 5,},
	}

	for _,v := range tests{
		f := CenteredAvg(v.data)
		if f != v.answer{
			t.Error("got:",f,"want:",v.answer)
		}
	}
}

//func TestCenteredAvg(t *testing.T) {
//	x := []int{1, 4, 6, 8, 100}
//	xs := CenteredAvg(x)
//	if xs != 6{
//		t.Error("Expected:",6,"\nGot:",xs)
//	}
//}

func ExampleCenteredAvg() {
	x := []int{1, 4, 6, 8, 100}
	fmt.Println(CenteredAvg(x))
	// Output:
	// 6
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i:=0;i<b.N;i++{
		mymath.Sum(1, 4, 6, 8, 100)
	}
}
