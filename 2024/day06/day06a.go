package day06

import (
	"github.com/tmeire/adventofcode/io"
)

func printGrid(grid [][]byte) {
	var b byte
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == b {
				print(string('.'))
			} else {
				print(string('X'))
			}
		}
		println()
	}
}

const (
	UP    byte = 0b0001
	RIGHT byte = 0b0010
	DOWN  byte = 0b0100
	LEFT  byte = 0b1000
)

func next(x, y int, dir byte) (int, int) {
	switch dir {
	case UP:
		return x - 1, y
	case RIGHT:
		return x, y + 1
	case DOWN:
		return x + 1, y
	case LEFT:
		return x, y - 1
	}
	panic("unknown direction")
}

func isObstacle(grid [][]byte, x, y int) bool {
	if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) {
		return grid[x][y] == '#' || grid[x][y] == '0'
	}
	return false
}

func makeGrid(h, w int) [][]byte {
	var grid [][]byte
	for i := 0; i < h; i++ {
		grid = append(grid, make([]byte, w))
	}
	return grid
}

func loops(grid [][]byte, x, y int, dir byte) bool {
	mirror := makeGrid(len(grid), len(grid[0]))

	//defer printGrid(mirror)

	for 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) {
		if mirror[x][y]&dir != 0 {
			return true
		}
		nextX, nextY := next(x, y, dir)
		if isObstacle(grid, nextX, nextY) {
			// turn right
			dir = (dir << 1) % 0b1111
		} else {
			// mark the position as visited
			mirror[x][y] = mirror[x][y] | dir
			x, y = nextX, nextY
		}
	}
	return false
}

func pathLen(grid [][]byte, x, y int, dir byte) int {
	mirror := makeGrid(len(grid), len(grid[0]))
	for 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) {
		nextX, nextY := next(x, y, dir)
		if isObstacle(grid, nextX, nextY) {
			// turn right
			dir = (dir << 1) % 0b1111
		} else {
			// mark the position as visited
			mirror[x][y] = dir
			x, y = nextX, nextY
		}
	}

	printGrid(grid)

	// Count all the visited positions
	l := 0
	for _, row := range mirror {
		for _, c := range row {
			if c != 0 {
				l++
			}
		}
	}
	return l
}

func Solve() {
	grid, err := io.ReadByteLinesFromFile("./2024/day06/input.txt")
	if err != nil {
		panic(err)
	}

	// Find the starting position
	var x, y int
	for i, row := range grid {
		for j, c := range row {
			if c == '^' {
				x = i
				y = j
			}
		}
	}

	l := pathLen(grid, x, y, UP)
	println(l)

	// TODO: we can trim down the search space to the positions the guard has visited in the first part
	positions := 0
	for i, row := range grid {
		for j, c := range row {
			if c != '#' && c != '^' {
				grid[i][j] = '0'
				if loops(grid, x, y, UP) {
					positions++
				}
				grid[i][j] = '.'
			}
		}
	}
	println(positions)
}
