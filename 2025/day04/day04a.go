package day04

import (
	"fmt"
	"io"
	"os"
)

func Solve() {
	grid := read("./2025/day04/input.txt")

	var t int
	for {
		var n int
		for i, row := range grid {
			for j, v := range row {
				if v == 0 {
					continue
				}

				if grid[i-1][j-1]+grid[i-1][j]+grid[i-1][j+1]+
					grid[i][j-1]+grid[i][j+1]+
					grid[i+1][j-1]+grid[i+1][j]+grid[i+1][j+1] < 4 {
					grid[i][j] = 0
					n++
				}
			}
		}
		if n == 0 {
			break
		}
		t += n
	}
	println(t)
}

func read(fn string) [][]int {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var grid [][]int
	for {
		row := []int{0}
		var j byte
		for {
			_, err := fmt.Fscanf(f, "%c", &j)
			if err != nil {
				if err == io.EOF {
					row = append(row, 0)
					empty := make([]int, len(row))
					return append([][]int{empty}, append(grid, row, empty)...)
				}
				panic(err.Error())
			}
			if j == '\n' {
				break
			}
			var v int
			if j == '@' {
				v = 1
			}
			row = append(row, v)
		}
		row = append(row, 0)
		grid = append(grid, row)
	}
}
