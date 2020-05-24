package mystr

import (
	"strings"
	"testing"
)



func TestCat(t *testing.T){
	s := "Shaken not stirred"
	xs := strings.Split(s," ")
	s = Cat(xs)

	if s != "Shaken not stirred"{
		t.Error("\nExpected Concate:","Shaken not stirred","\ngot:",s)
	}
}

func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s," ")
	s = Join(xs)

	if s != "Shaken not stirred"{
		t.Error("Expected Join: ","Shaken not stirred"," got:",s)
	}
}

const s = "We ask ourselves, who am i to be brilliant"
var xs []string

func BenchmarkCat(b *testing.B){
	for i:= 0; i< b.N; i++{
		n_s := ""
		xs = strings.Split(s,"")
		for _,v :=range xs{
			n_s += v
			n_s += " "
 		}
	}
}

func BenchmarkJoin(b *testing.B){
	for i:= 0; i< b.N; i++{
		xs := strings.Split(s," ")
		strings.Join(xs, " ")
	}
}