package day14

import (
	"fmt"
	"math/bits"

	"github.com/tmeire/adventofcode/algo/knothash"
)

func htoi(b byte) byte {
	if '0' <= b && b <= '9' {
		return b - '0'
	}
	return 10 + (b - 'a')
}

func set(grid [][]int, mask [][]bool, i, j, c int) {
	// skip if this position is not occupied
	if !mask[i][j] {
		return
	}
	// skip if this position is already colored
	if grid[i][j] != 0 {
		return
	}

	grid[i][j] = c

	if i > 0 {
		set(grid, mask, i-1, j, c)
	}
	if i < len(grid)-1 {
		set(grid, mask, i+1, j, c)
	}
	if j > 0 {
		set(grid, mask, i, j-1, c)
	}
	if j < len(grid)-1 {
		set(grid, mask, i, j+1, c)
	}
}

func color(grid [][]int, mask [][]bool) int {
	c := 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 && mask[i][j] {
				set(grid, mask, i, j, c)
				c++
			}
		}
	}
	return c - 1
}

func Solve() {
	input := "jzgqcdpd"
	//input := "flqrgnkx"

	mask := make([][]bool, 128)
	for i := 0; i < 128; i++ {
		mask[i] = make([]bool, 128)
	}

	bitcount := 0
	for i := 0; i < 128; i++ {
		hex := knothash.Hash(fmt.Sprintf("%s-%d", input, i))

		for j := range hex {
			u := uint8(htoi(hex[j]))

			// Keep track of the bitcount for part A
			bitcount += bits.OnesCount8(u)

			// Create the mask for part B
			for a := 0; a < 4; a++ {
				m := uint8(1 << uint(a))
				mask[i][(j+1)*4-a-1] = (u&m == m)
			}
		}
	}
	fmt.Println("Part A", bitcount)

	grid := make([][]int, 128)
	for i := 0; i < 128; i++ {
		grid[i] = make([]int, 128)
	}
	fmt.Println("Part B", color(grid, mask))
}
