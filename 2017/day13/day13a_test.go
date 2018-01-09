package day13

import "testing"

func BenchmarkPartA(b *testing.B) {
	fw := Firewall([]Layer{
		&SecurityLayer{0, 4, 0, true},
		&SecurityLayer{1, 2, 0, true},
		&SecurityLayer{2, 3, 0, true},
		&ZeroLayer{},
		&SecurityLayer{4, 5, 0, true},
		&ZeroLayer{},
		&SecurityLayer{6, 8, 0, true},
		&ZeroLayer{},
		&SecurityLayer{8, 6, 0, true},
		&ZeroLayer{},
	})

	for i := 0; i < b.N; i++ {
		fw.Caught(10000)
	}
}
