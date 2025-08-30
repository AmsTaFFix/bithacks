package main

import "testing"

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := bool2int1(true)
		_ = a
	}
}
func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := bool2int2(true)
		_ = a
	}
}
