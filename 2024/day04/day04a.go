package day04

import (
	"github.com/tmeire/adventofcode/io"
)

func pad(grid [][]byte) [][]byte {
	buf := make([]byte, len(grid[0]))
	grid = append(grid, buf, buf, buf)
	grid = append([][]byte{buf, buf, buf}, grid...)

	var pg [][]byte
	for i := 0; i < len(grid); i++ {
		pl := make([]byte, len(grid[i])+6)
		copy(pl[3:], grid[i])
		pg = append(pg, pl)
	}
	return pg
}

func printGrid(grid [][]byte) {
	var b byte
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == b {
				print(string('-'))
			} else {
				print(string(grid[i][j]))
			}
		}
		println()
	}
}

func isXMAS(a, b, c, d byte) bool {
	return a == 'X' && b == 'M' && c == 'A' && d == 'S'
}

func Solve() {
	grid, err := io.ReadByteLinesFromFile("./2024/day04/input.txt")
	if err != nil {
		panic(err)
	}

	// I don't want to do bounds checks
	grid = pad(grid)

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 'X' {
				continue
			}

			// Look right
			if isXMAS(grid[i][j], grid[i][j+1], grid[i][j+2], grid[i][j+3]) {
				count++
			}
			// Look left
			if isXMAS(grid[i][j], grid[i][j-1], grid[i][j-2], grid[i][j-3]) {
				count++
			}
			// Look down
			if isXMAS(grid[i][j], grid[i+1][j], grid[i+2][j], grid[i+3][j]) {
				count++
			}
			// Look up
			if isXMAS(grid[i][j], grid[i-1][j], grid[i-2][j], grid[i-3][j]) {
				count++
			}

			// Bottom right
			if isXMAS(grid[i][j], grid[i+1][j+1], grid[i+2][j+2], grid[i+3][j+3]) {
				count++
			}
			// Look top left
			if isXMAS(grid[i][j], grid[i-1][j-1], grid[i-2][j-2], grid[i-3][j-3]) {
				count++
			}
			// Bottom left
			if isXMAS(grid[i][j], grid[i+1][j-1], grid[i+2][j-2], grid[i+3][j-3]) {
				count++
			}
			// Top Right
			if isXMAS(grid[i][j], grid[i-1][j+1], grid[i-2][j+2], grid[i-3][j+3]) {
				count++
			}
		}
	}
	println(count)

	xmas := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 'A' {
				continue
			}

			if isX_MAS(grid[i][j], grid[i-1][j-1], grid[i-1][j+1], grid[i+1][j+1], grid[i+1][j-1]) {
				xmas++
			}
		}
	}
	println(xmas)
}

/**
 *  b . c
 *  . a .
 *  e . d
 */
func isX_MAS(a, b, c, d, e byte) bool {
	return a == 'A' &&
		((b == 'M' && d == 'S') || (b == 'S' && d == 'M')) &&
		((c == 'M' && e == 'S') || (c == 'S' && e == 'M'))
}
