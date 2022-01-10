package main

import "testing"

func BenchmarkUnoptimized(b *testing.B) {
	unoptimized := new(Unoptimized)
	d := data{make([]bool, 100)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		unoptimized.toggle(&d)
	}
}

func BenchmarkOptimized1(b *testing.B) {
	optimized1 := new(Optimized1)
	d := data{make([]bool, 100)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		optimized1.toggle(&d)
	}
}

func BenchmarkOptimized2(b *testing.B) {
	optimized2 := new(Optimized2)
	d := data{make([]bool, 100)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		optimized2.toggle(&d)
	}
}
