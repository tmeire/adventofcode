package day08

import (
	"os"
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	os.Args[1] = "/Users/thomas.meire/Source/gosrc/src/github.com/tmeire/adventofcode/2017/day08/input.txt"
	for i := 0; i < b.N; i++ {
		Solve()
	}
}

func BenchmarkSolve_v2(b *testing.B) {
	os.Args[1] = "/Users/thomas.meire/Source/gosrc/src/github.com/tmeire/adventofcode/2017/day08/input.txt"
	for i := 0; i < b.N; i++ {
		Solve_v2()
	}
}
