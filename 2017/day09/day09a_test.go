package day09

import (
	"os"
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	os.Args[1] = "/Users/thomas.meire/Source/gosrc/src/github.com/tmeire/adventofcode/2017/day09/input.txt"
	for i := 0; i < b.N; i++ {
		Solve()
	}
}
