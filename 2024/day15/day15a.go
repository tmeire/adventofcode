package day15

import (
	"github.com/tmeire/adventofcode/io"
)

type grid [][]byte

func next(x, y int, dir byte) (int, int) {
	switch dir {
	case '^':
		return x - 1, y
	case 'v':
		return x + 1, y
	case '>':
		return x, y + 1
	case '<':
		return x, y - 1
	default:
		panic("unknown dir " + string(dir) + " on map")
	}
}

func (g grid) move(x, y int, dir byte) (int, int) {
	nx, ny := next(x, y, dir)
	if g.isFree(nx, ny, dir) {
		g[nx][ny] = g[x][y]
		g[x][y] = '.'
		return nx, ny
	} else {
		return x, y
	}
}

// Only to be called when dir is up or down
func (g grid) moveWithoutCheck(x, y int, dir byte) {
	switch g[x][y] {
	case '.':
		// it's possible that this field was already cleared by moving the other part of the box
		return
	case '[':
		nxl, nyl := next(x, y, dir)
		nxr, nyr := next(x, y+1, dir)

		g.moveWithoutCheck(nxl, nyl, dir)
		g.moveWithoutCheck(nxr, nyr, dir)

		g[nxl][nyl] = g[x][y]
		g[nxr][nyr] = g[x][y+1]
		g[x][y] = '.'
		g[x][y+1] = '.'

	case ']':
		nxl, nyl := next(x, y-1, dir)
		nxr, nyr := next(x, y, dir)

		g.moveWithoutCheck(nxl, nyl, dir)
		g.moveWithoutCheck(nxr, nyr, dir)

		g[nxl][nyl] = g[x][y-1]
		g[nxr][nyr] = g[x][y]
		g[x][y-1] = '.'
		g[x][y] = '.'
	}
}

func (g grid) canBeFree(x, y int, dir byte) bool {
	switch g[x][y] {
	case '#':
		// Walls aren'r moveable, so always false
		return false
	case '.':
		// Dots are free
		return true
	case 'O':
		// It's a small box, let's see if we can move it
		nx, ny := next(x, y, dir)
		return g.canBeFree(nx, ny, dir)
	case '[':
		// also check and move the ]
		switch dir {
		case '^':
			nxl, nyl := next(x, y, dir)
			nxr, nyr := next(x, y+1, dir)

			return g.canBeFree(nxl, nyl, dir) && g.canBeFree(nxr, nyr, dir)
		case 'v':
			nxl, nyl := next(x, y, dir)
			nxr, nyr := next(x, y+1, dir)

			return g.canBeFree(nxl, nyl, dir) && g.canBeFree(nxr, nyr, dir)
		case '>':
			return g.canBeFree(x, y+1, dir)
		case '<':
			return g.canBeFree(x, y-1, dir)
		default:
			panic("unknown dir " + string(dir) + " on map")
		}
	case ']':
		// also check and move the [
		switch dir {
		case '^':
			nxl, nyl := next(x, y-1, dir)
			nxr, nyr := next(x, y, dir)

			return g.canBeFree(nxl, nyl, dir) && g.canBeFree(nxr, nyr, dir)
		case 'v':
			nxl, nyl := next(x, y-1, dir)
			nxr, nyr := next(x, y, dir)

			return g.canBeFree(nxl, nyl, dir) && g.canBeFree(nxr, nyr, dir)
		case '>':
			return g.canBeFree(x, y+1, dir)
		case '<':
			return g.canBeFree(x, y-1, dir)
		default:
			panic("unknown dir " + string(dir) + " on map")
		}
	default:
		panic("unknown char " + string(g[x][y]) + " on map")
	}
}

func (g grid) isFree(x, y int, dir byte) bool {
	switch g[x][y] {
	case '#':
		// Walls aren'r moveable, so always false
		return false
	case '.':
		// Dots are free
		return true
	case 'O':
		// It's a small box, let's see if we can move it
		nx, ny := next(x, y, dir)
		if g.isFree(nx, ny, dir) {
			g[nx][ny] = g[x][y]
			g[x][y] = '.'
			return true
		} else {
			return false
		}
	case '[':
		// also check and move the ]
		nxl, nyl := next(x, y, dir)
		nxr, nyr := next(x, y+1, dir)
		switch dir {
		case '^':
			if g.canBeFree(nxl, nyl, dir) && g.canBeFree(nxr, nyr, dir) {
				g.moveWithoutCheck(nxl, nyl, dir)
				g.moveWithoutCheck(nxr, nyr, dir)
				g[nxl][nyl] = g[x][y]
				g[nxr][nyr] = g[x][y+1]
				g[x][y] = '.'
				g[x][y+1] = '.'
				return true
			} else {
				return false
			}
		case 'v':
			if g.canBeFree(nxl, nyl, dir) && g.canBeFree(nxr, nyr, dir) {
				g.moveWithoutCheck(nxl, nyl, dir)
				g.moveWithoutCheck(nxr, nyr, dir)
				g[nxl][nyl] = g[x][y]
				g[nxr][nyr] = g[x][y+1]
				g[x][y] = '.'
				g[x][y+1] = '.'
				return true
			} else {
				return false
			}
		case '>':
			if g.isFree(x, y+1, dir) {
				g[x][y+1] = '['
				g[x][y] = '.'
				return true
			} else {
				return false
			}
		case '<':
			if g.isFree(x, y-1, dir) {
				g[x][y-1] = '['
				g[x][y] = '.'
				return true
			} else {
				return false
			}
		default:
			panic("unknown dir " + string(dir) + " on map")
		}
	case ']':
		// also check and move the one at j-1
		nxl, nyl := next(x, y-1, dir)
		nxr, nyr := next(x, y, dir)

		switch dir {
		case '^':
			if g.canBeFree(nxl, nyl, dir) && g.canBeFree(nxr, nyr, dir) {
				g.moveWithoutCheck(nxl, nyl, dir)
				g.moveWithoutCheck(nxr, nyr, dir)
				g[nxl][nyl] = g[x][y-1]
				g[nxr][nyr] = g[x][y]
				g[x][y-1] = '.'
				g[x][y] = '.'
				return true
			} else {
				return false
			}
		case 'v':
			if g.canBeFree(nxl, nyl, dir) && g.canBeFree(nxr, nyr, dir) {
				g.moveWithoutCheck(nxl, nyl, dir)
				g.moveWithoutCheck(nxr, nyr, dir)
				g[nxl][nyl] = g[x][y-1]
				g[nxr][nyr] = g[x][y]
				g[x][y-1] = '.'
				g[x][y] = '.'
				return true
			} else {
				return false
			}
		case '>':
			if g.isFree(x, y+1, dir) {
				g[x][y+1] = ']'
				g[x][y] = '.'
				return true
			} else {
				return false
			}
		case '<':
			if g.isFree(x, y-1, dir) {
				g[x][y-1] = ']'
				g[x][y] = '.'
				return true
			} else {
				return false
			}
		default:
			panic("unknown dir " + string(dir) + " on map")
		}
	default:
		panic("unknown char " + string(g[x][y]) + " on map")
	}
}

func (g grid) scale() grid {
	var newGrid [][]byte
	for i := 0; i < len(g); i++ {
		var newLine []byte
		for j := 0; j < len(g[i]); j++ {
			switch g[i][j] {
			case '@':
				newLine = append(newLine, '@', '.')
			case 'O':
				newLine = append(newLine, '[', ']')
			default:
				newLine = append(newLine, g[i][j], g[i][j])
			}
		}
		newGrid = append(newGrid, newLine)
	}
	return newGrid
}

func Solve() {
	input, err := io.ReadByteLinesFromFile("./2024/day15/input.txt")
	if err != nil {
		panic(err)
	}

	g := grid(input[:len(input)-2])

	g = g.scale()
	printGrid(g)

	// Find the starting position of the robot
	var rx, ry int
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == '@' {
				rx, ry = i, j
			}
		}
	}

	// Run all its moves
	for _, dir := range input[len(input)-1] {
		rx, ry = g.move(rx, ry, dir)
		//printGrid(g)
	}

	// Get all the coordinates
	var sum int
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == '[' {
				x1 := i
				x2 := len(g[i]) - (i + 1)
				if x1 < x2 {
					sum += x1*100 + j
				} else {
					sum += x2*100 + j
				}
			} else if g[i][j] == '0' {
				sum = i*100 + j
			}
		}
	}
	println(sum)
}

func printGrid(grid [][]byte) {
	var b byte
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == b {
				print(string('.'))
			} else {
				print(string(grid[i][j]))
			}
		}
		println()
	}
}
