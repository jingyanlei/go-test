package test

import (
	"testing"
)

func BenchmarkTest2(b *testing.B)  {
	//fmt.Println("b.Ns:",b.N)
	//b.SetBytes(1)
	for i:=0; i<b.N; i++ {
		test2()
	}
}

func BenchmarkTest3(b *testing.B)  {
	//fmt.Println("b.Ns:",b.N)
	//b.SetBytes(1)
	for i:=0; i<b.N; i++ {
		test3()
	}
}

