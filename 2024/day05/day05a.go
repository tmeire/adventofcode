package day05

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

func parseRule(l string) (a, b int) {
	p := strings.Split(l, "|")
	return atoi(p[0]), atoi(p[1])
}

func parseUpdate(l string) []int {
	ps := strings.Split(l, ",")
	var update []int
	for _, p := range ps {
		update = append(update, atoi(p))
	}
	return update
}

func parse(lines []string) (map[int][]int, [][]int) {
	rulesMap := make(map[int][]int)
	updates := make([][]int, 0)

	rules := true
	for _, l := range lines {
		if l == "" {
			rules = false
			continue
		}
		if rules {
			a, b := parseRule(l)
			rulesMap[a] = append(rulesMap[a], b)
		} else {
			updates = append(updates, parseUpdate(l))
		}
	}

	return rulesMap, updates
}

func isCorrectlyOrdered(update []int, rules map[int][]int) bool {
	for i, p := range update {
		for _, r := range rules[p] {
			for j := 0; j < i; j++ {
				if update[j] == r {
					return false
				}
			}
		}
	}
	return true
}

func sort(update []int, rules map[int][]int) []int {
	slices.SortFunc(update, func(a, b int) int {
		for _, r := range rules[a] {
			if b == r {
				return -1
			}
		}
		for _, r := range rules[b] {
			if a == r {
				return 1
			}
		}
		return 0
	})
	return update
}

func Solve() {
	lines, err := io.ReadLinesFromFile("./2024/day05/input.txt")
	if err != nil {
		panic(err)
	}

	rules, updates := parse(lines)

	correctSum := 0
	correctedSum := 0
	for _, update := range updates {
		if isCorrectlyOrdered(update, rules) {
			correctSum += update[len(update)/2]
		} else {
			corrected := sort(update, rules)
			correctedSum += corrected[len(update)/2]
		}
	}
	println(correctSum)
	println(correctedSum)
}
