package day03

import "fmt"

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func location(input int) (int, int) {
	edge := 1
	size := 1
	for size < input {
		edge += 2
		size = edge * edge
	}
	mid := edge/2 + 1

	prevedge := edge - 2
	prevsize := prevedge * prevedge

	diff := input - prevsize

	q := (size - prevsize) / 4

	var x, y int
	if diff < q { // right edge
		x = edge
		y = edge - diff
	} else if diff < 2*q { // upper edge
		x = edge - (diff - q)
		y = 1
	} else if diff < 3*q { // left edge
		x = 1
		y = diff - 2*q + 1
	} else { //if diff < 4*q { // lower edge
		x = diff - 3*q + 1
		y = edge
	}
	// recenter x and y around 0
	return x - mid, y - mid
}

type grid struct {
	g map[int]map[int]int
}

func (g *grid) put(x, y int, sum int) {
	r, ok := g.g[x]
	if !ok {
		r = make(map[int]int)
		g.g[x] = r
	}
	r[y] = sum
}

func (g *grid) value(x, y int) int {
	if r, ok := g.g[x]; ok {
		if c, ok := r[y]; ok {
			return c
		}
	}
	return 0
}

func (g *grid) compute(x, y int) int {
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			sum += g.value(x+i, y+j)
		}
	}
	g.put(x, y, sum)
	return sum
}

func partA(input int) {
	x, y := location(input)

	fmt.Println("Part A", abs(x)+abs(y))
}

func partB(input int) {
	g := grid{make(map[int]map[int]int)}
	g.put(0, 0, 1)

	v := 0
	for i := 2; i <= input; i++ {
		v = g.compute(location(i))
		if v > input {
			break
		}
	}

	fmt.Println("Part B", v)
}

func Solve() {
	input := 312051
	partA(input)
	partB(input)
}
