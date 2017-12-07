package day06

import "fmt"

type banks []int

func (b banks) max() (int, int) {
	maxIdx := 0
	max := b[0]
	for i := 1; i < len(b); i++ {
		if b[i] > max {
			max = b[i]
			maxIdx = i
		}
	}
	return maxIdx, max
}

func (b banks) redistribute() {
	idx, val := b.max()

	b[idx] = 0
	idx = (idx + 1) % len(b)

	for ; val > 0; val-- {
		b[idx]++
		idx = (idx + 1) % len(b)
	}
}

func (b banks) String() string {
	return fmt.Sprintf("%v", []int(b))
}

func solve(b banks) {
	steps := 0

	set := make(map[string]int)

	seen := func(s string) bool {
		_, ok := set[s]
		return ok
	}

	for !seen(b.String()) {
		set[b.String()] = steps
		b.redistribute()
		steps++
	}

	fmt.Println("Part A", steps)
	fmt.Println("Part B", steps-set[b.String()])
}

func Solve() {
	input := []int{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}

	solve(banks(input))
}
