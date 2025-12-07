package day07

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

func propagation(lines [][]byte, i, j int) int {
	if i >= len(lines) {
		return 0
	}

	b := lines[i][j]
	switch b {
	case '.':
		lines[i][j] = '|'
		return propagation(lines, i+1, j)
	case 'S':
		lines[i][j] = '|'
		return propagation(lines, i+1, j)
	case '^':
		n := 1
		if j > 0 {
			n += propagation(lines, i, j-1)
		}
		if j < len(lines[i])-1 {
			n += propagation(lines, i, j+1)
		}
		return n
	}
	return 0
}

func timelines(lines [][]byte, colored [][]int, i, j int) int {
	if i >= len(lines) {
		return 1
	}
	if colored[i][j] > 0 {
		return colored[i][j]
	}

	b := lines[i][j]
	switch b {
	case '.':
		n := timelines(lines, colored, i+1, j)
		colored[i][j] = n
		return n
	case 'S':
		n := timelines(lines, colored, i+1, j)
		colored[i][j] = n
		return n
	case '^':
		n := 0
		if j > 0 {
			n += timelines(lines, colored, i, j-1)
		}
		if j < len(lines[i])-1 {
			n += timelines(lines, colored, i, j+1)
		}
		colored[i][j] = n
		return n
	}
	panic("unreachable")
}

func zero(lines [][]byte) [][]int {
	var z [][]int
	for i := range lines {
		z = append(z, make([]int, len(lines[i])))
	}
	return z
}

func Solve() {
	lines := read("2025/day07/input.txt")

	colored := zero(lines)

	for i, b := range lines[0] {
		if b == 'S' {
			println(timelines(lines, colored, 0, i))
		}
	}
	for i, b := range lines[0] {
		if b == 'S' {
			println(propagation(lines, 0, i))
		}
	}
	for _, b := range lines {
		for _, c := range b {
			print(string(c))
		}
		println()
	}
}

func read(fn string) [][]byte {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fb := bufio.NewReader(f)
	var lines [][]byte
	for {
		l, _, err := fb.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return lines
			}
			panic(err.Error())
		}
		lines = append(lines, bytes.Clone(l))
	}
	return lines
}
