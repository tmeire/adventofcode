package day10

import "testing"

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve()
	}
}
