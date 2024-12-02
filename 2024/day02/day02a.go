package day02

import (
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

func convert(report string) []int {
	levelstrings := strings.Split(report, " ")

	var levels []int
	for _, l := range levelstrings {
		levels = append(levels, atoi(l))
	}
	return levels
}

func isTolerated(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		var sub []int
		if i == 0 {
			sub = levels[1:]
		} else {
			sub = make([]int, i)
			copy(sub, levels[:i])
			sub = append(sub, levels[i+1:]...)
		}
		if isSafe(sub) {
			return true
		}
	}
	return false
}

func absdiff(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}
	return b - a, false
}

func isSafe(levels []int) bool {
	mustDesc := levels[0] > levels[1]
	for i := 1; i < len(levels); i++ {
		diff, desc := absdiff(levels[i-1], levels[i])
		if mustDesc != desc || diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func Solve() {
	reports, err := io.ReadLinesFromFile("./2024/day02/input.txt")
	if err != nil {
		panic(err)
	}

	// part 1 & 2
	safe := 0
	tolerated := 0
	for _, report := range reports {
		levels := convert(report)
		if isSafe(levels) {
			safe++
		}
		if isTolerated(levels) {
			tolerated++
		}
	}
	println(safe)
	println(tolerated)
}
