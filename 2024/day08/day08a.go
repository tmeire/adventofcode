package day08

import (
	"github.com/tmeire/adventofcode/io"
)

func makeGrid(h, w int) [][]byte {
	var grid [][]byte
	for i := 0; i < h; i++ {
		grid = append(grid, make([]byte, w))
	}
	return grid
}

type tuple struct {
	x, y int
}

func calc(a, b tuple) tuple {
	return tuple{
		2*a.x - b.x,
		2*a.y - b.y,
	}
}

func calcP2(w, h int, a, b tuple) []tuple {
	var results []tuple
	i := 0
	for {
		t := tuple{
			a.x + i*(a.x-b.x),
			a.y + i*(a.y-b.y),
		}
		if t.x < 0 || t.x >= w || t.y < 0 || t.y >= h {
			return results
		}
		results = append(results, t)
		i++
	}
}

func printGrid(grid [][]byte) {
	var b byte
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == b {
				print(string('.'))
			} else {
				print(string('#'))
			}
		}
		println()
	}
}

func Solve() {
	grid, err := io.ReadByteLinesFromFile("./2024/day08/input.txt")
	if err != nil {
		panic(err)
	}

	antennas := make(map[byte][]tuple)
	for i, row := range grid {
		for j, f := range row {
			if f != '.' {
				antennas[f] = append(antennas[f], tuple{i, j})
			}
		}
	}

	mirror := makeGrid(len(grid), len(grid[0]))

	// compute all the antinodes
	for _, nodes := range antennas {
		for i, node1 := range nodes {
			for j, node2 := range nodes {
				if i == j {
					continue
				}
				an := calc(node1, node2)
				if 0 <= an.x && an.x < len(mirror) && 0 <= an.y && an.y < len(mirror[0]) {
					mirror[an.x][an.y] = '#'
				}
			}
		}
	}

	printGrid(mirror)

	antinodes := 0
	for _, row := range mirror {
		for _, f := range row {
			if f == '#' {
				antinodes++
			}
		}
	}
	println(antinodes)

	mirrorP2 := makeGrid(len(grid), len(grid[0]))

	// compute all the antinodes
	for _, nodes := range antennas {
		for i, node1 := range nodes {
			for j, node2 := range nodes {
				if i == j {
					continue
				}
				ans := calcP2(len(grid), len(grid[0]), node1, node2)
				println(len(ans))
				for _, an := range ans {
					mirrorP2[an.x][an.y] = '#'
				}
			}
		}
	}
	printGrid(mirrorP2)

	antinodesP2 := 0
	for _, row := range mirrorP2 {
		for _, f := range row {
			if f == '#' {
				antinodesP2++
			}
		}
	}
	println(antinodesP2)

}
