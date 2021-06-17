package main

import (
	"testing"
)

func BenchmarkF(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f()
	}
}

func f() {
	var slice []int

	for i := 0; i < 100000; i++ {
		slice = append(slice, i)
	}
}
