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

func isSafeWithDirection(levels []int) (bool, bool) {
	mustDesc := levels[0] > levels[1]
	for i := 1; i < len(levels); i++ {
		diff, desc := absdiff(levels[i-1], levels[i])
		if mustDesc != desc || diff < 1 || diff > 3 {
			return false, mustDesc
		}
	}
	return true, mustDesc
}

func isSafe(levels []int) bool {
	ok, _ := isSafeWithDirection(levels)
	return ok
}

func isSafeWithoutCopy(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	if isSafe(levels[1:]) {
		return true
	}

	if isSafe(levels[:len(levels)-1]) {
		return true
	}

	for i := 1; i < len(levels)-1; i++ {
		diff, desc := absdiff(levels[i-1], levels[i+1])
		switch i {
		case 1:
			ok, mustDesc := isSafeWithDirection(levels[i+1:])
			if ok && mustDesc == desc && diff >= 1 && diff <= 3 {
				return true
			}
		case len(levels) - 2:
			ok, mustDesc := isSafeWithDirection(levels[:i])
			if ok && mustDesc == desc && diff >= 1 && diff <= 3 {
				return true
			}
		default:
			ok1, mustDesc1 := isSafeWithDirection(levels[:i])
			ok2, mustDesc2 := isSafeWithDirection(levels[i+1:])
			if ok1 && ok2 && mustDesc1 == mustDesc2 && mustDesc1 == desc && diff >= 1 && diff <= 3 {
				return true
			}
		}
	}

	return false
}

func isGoodDiff(diff int) bool {
	return 1 <= diff && diff <= 3
}

func bigger(a, b int) int {
	if a < b {
		return 1
	}
	return -1
}

func isAscending(levels []int) int {
	ascCount := 0
	for i := 0; i < len(levels)-1; i++ {
		ascCount += bigger(levels[i], levels[i+1])
	}
	if ascCount > 0 {
		return -1
	}
	return 1
}

func isSafeOnV2(levels []int) bool {
	asc := isAscending(levels)
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i] - levels[i+1]
		if !(isGoodDiff(asc * diff)) {
			// We're at the end, we can ignore the last entry
			if i+2 == len(levels) {
				return true
			}
			// try to see if it works by ignoring i+1
			diff1 := levels[i] - levels[i+2]
			if isGoodDiff(asc * diff1) {
				// it's ok to jump from i to i+2, check if the rest of the slice is ok
				return isSafeOnV2Sub(levels[i+2:], asc)
			}
			if i == 0 {
				// if we're at the start of the slice, but can't skip 1, then try to skip 0
				return isSafeOnV2Sub(levels[1:], asc)
			}
			diff0 := levels[i-1] - levels[i+1]
			if isGoodDiff(asc * diff0) {
				// it's ok to jump from i-1 to i+1, check if the rest of the slice is ok
				return isSafeOnV2Sub(levels[i+1:], asc)
			}
			return false
		}
	}
	return true
}

func isSafeOnV2Sub(levels []int, asc int) bool {
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i] - levels[i+1]
		if !(isGoodDiff(asc * diff)) {
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
	toleratedV2 := 0
	toleratedOn := 0
	for _, report := range reports {
		levels := convert(report)
		if isSafe(levels) {
			safe++
		}
		if isTolerated(levels) {
			tolerated++
		}
		if isSafeWithoutCopy(levels) {
			toleratedV2++
		}
		if isSafeOnV2(levels) {
			toleratedOn++
		}
	}
	println(safe)
	println(tolerated)
	println(toleratedV2)
	println(toleratedOn)
}
