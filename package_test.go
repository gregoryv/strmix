package compstr

import (
	"fmt"
	"testing"
)

func Example_format() {
	fmt.Println(Split("ServeHTTP"))
	fmt.Println(Split("CallRPCMethod"))
	fmt.Println(Split("ExampleCar_Model"))
	// output:
	// [Serve HTTP]
	// [Call RPC Method]
	// [Example Car Model]
}

func BenchmarkSplit(b *testing.B) {
	in := "CallRPCHandler"
	for i := 0; i < b.N; i++ {
		_ = Split(in)
	}
}
func BenchmarkAppendWords(b *testing.B) {
	in := "CallRPCHandler"
	compstr := make([]string, 0, 3)
	for i := 0; i < b.N; i++ {
		_ = AppendWords(compstr, in)
	}
}

func ExampleSplit() {
	fmt.Println(Split("CallRPCHandler_nowOrLater"))
	fmt.Println(Split("H20Chemical"))
	fmt.Println(Split("hiddenField"))
	fmt.Println(Split("word"))
	fmt.Println(Split("ab__cd"))
	fmt.Println(Split("1_2_3A"))
	// output:
	// [Call RPC Handler now Or Later]
	// [H20 Chemical]
	// [hidden Field]
	// [word]
	// [ab cd]
	// [1 2 3A]
}
