package day10

import (
	"fmt"

	"github.com/tmeire/adventofcode/io"
)

func score(grid [][]byte, x, y int, step byte, tops map[string]struct{}) (int, int) {
	if step == '9' {
		h := fmt.Sprintf("%d-%d", x, y)
		if _, ok := tops[h]; ok {
			return 0, 1
		}
		tops[h] = struct{}{}
		return 1, 1
	}

	scores := 0
	ratings := 0
	if x > 0 && grid[x-1][y] == step+1 {
		s, r := score(grid, x-1, y, step+1, tops)
		scores += s
		ratings += r
	}
	if x < len(grid)-1 && grid[x+1][y] == step+1 {
		s, r := score(grid, x+1, y, step+1, tops)
		scores += s
		ratings += r
	}
	if y > 0 && grid[x][y-1] == step+1 {
		s, r := score(grid, x, y-1, step+1, tops)
		scores += s
		ratings += r
	}
	if y < len(grid[0])-1 && grid[x][y+1] == step+1 {
		s, r := score(grid, x, y+1, step+1, tops)
		scores += s
		ratings += r
	}

	return scores, ratings
}

func Solve() {
	grid, err := io.ReadByteLinesFromFile("./2024/day10/input.txt")
	if err != nil {
		panic(err)
	}

	scores := 0
	ratings := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' {
				s, r := score(grid, i, j, '0', make(map[string]struct{}))
				scores += s
				ratings += r
			}
		}
	}
	println(scores)
	println(ratings)
}
