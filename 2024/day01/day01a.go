package day01

import (
	"slices"
	"strconv"
	"strings"

	"github.com/tmeire/adventofcode/io"
)

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func absdiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Solve() {
	lines, err := io.ReadLinesFromFile("./2024/day01/input.txt")
	if err != nil {
		panic(err)
	}

	var left, right []int
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		left = append(left, atoi(parts[0]))
		right = append(right, atoi(parts[1]))
	}

	// Part 1
	slices.Sort(left)
	slices.Sort(right)

	dist := 0
	for i, v := range left {
		dist += absdiff(v, right[i])
	}
	println(dist)

	// Part 2
	counts := make(map[int]int)
	for _, v := range right {
		counts[v]++
	}

	simscore := 0
	for _, v := range left {
		simscore += v * counts[v]
	}
	println(simscore)
}
