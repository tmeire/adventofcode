package day10

import (
	"fmt"

	"github.com/blackskad/adventofcode/algo/knothash"
)

func Solve() {
	l := make(knothash.List, 256)
	for i := 0; i < 256; i++ {
		l[i] = byte(i)
	}

	lengths := []int{97, 167, 54, 178, 2, 11, 209, 174, 119, 248, 254, 0, 255, 1, 64, 190}

	l.Sparse(lengths, 1)

	// Multiply as ints to avoid byte overflow
	fmt.Println("Part A:", int(l[0])*int(l[1]))

	input := "97,167,54,178,2,11,209,174,119,248,254,0,255,1,64,190"

	fmt.Printf("Part B: %s\n", knothash.Hash(input))
}
