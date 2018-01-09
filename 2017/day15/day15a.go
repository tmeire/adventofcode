package day15

import "fmt"

const (
	bitmask = uint64(0xFFFF)
)

type Generate func() uint64

func generator(start, factor, mod uint64) Generate {
	value := start
	return func() uint64 {
		value = (value * factor) % 2147483647
		if mod > 0 {
			for value%mod != 0 {
				value = (value * factor) % 2147483647
			}
		}
		return value
	}
}

func judge(a, b Generate, pairs int) int {
	matches := 0
	for i := 0; i < pairs; i++ {
		if a()&bitmask == b()&bitmask {
			matches++
		}
	}
	return matches
}

func Solve() {
	genA := generator(516, 16807, 0)
	genB := generator(190, 48271, 0)

	fmt.Println("Part A:", judge(genA, genB, 40000000))

	genA = pickyGenerator(516, 16807, 4)
	genB = pickyGenerator(190, 48271, 8)

	fmt.Println("Part B:", judge(genA, genB, 5000000))
}
