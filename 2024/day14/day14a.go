package day14

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/tmeire/adventofcode/io"
)

func atoi(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

type pair struct {
	x, y int
}

type robot struct {
	p pair
	v pair
}

var format = regexp.MustCompile(`p=([0-9]+),([0-9]+) v=(\-?[0-9]+),(\-?[0-9]+)`)

func read() []robot {
	lines, err := io.ReadLinesFromFile("./2024/day14/input.txt")
	if err != nil {
		panic(err)
	}

	var robots []robot
	for _, l := range lines {
		d := format.FindAllStringSubmatch(l, -1)
		fmt.Println(d)
		robots = append(robots, robot{
			p: pair{atoi(d[0][1]), atoi(d[0][2])},
			v: pair{atoi(d[0][3]), atoi(d[0][4])},
		})
	}
	return robots

}

func makeGrid(h, w int) [][]byte {
	var grid [][]byte
	for i := 0; i < h; i++ {
		grid = append(grid, make([]byte, w))
	}
	return grid
}
func printGrid(grid [][]byte) {
	var b byte
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == b {
				print(string('.'))
			} else {
				print(grid[i][j])
			}
		}
		println()
	}
}

func (r robot) move(w, h, s int) (int, int) {
	nx := (r.p.x + s*r.v.x) % w
	for nx < 0 {
		nx += w
	}
	ny := (r.p.y + s*r.v.y) % h
	for ny < 0 {
		ny += h
	}
	return nx, ny
}

func Solve() {
	w := 101
	h := 103
	s := 100

	//p=2,4 v=2,-3
	r := robot{p: pair{2, 4}, v: pair{2, -3}}
	println(r.move(w, h, 1))
	println(r.move(w, h, 2))
	println(r.move(w, h, 3))
	println(r.move(w, h, 4))
	println(r.move(w, h, 5))

	robots := read()

	grid := makeGrid(h, w)
	var lt, rt, rb, lb int
	for _, r := range robots {
		nx, ny := r.move(w, h, s)
		grid[ny][nx]++

		switch {
		case nx == w/2.:
			// ignore
		case nx < w/2.:
			// left-side
			switch {
			case ny == h/2.:
				// ignore
			case ny < h/2.:
				// top
				lt++
			case ny > h/2.:
				// bottom
				lb++
			}
		case nx > w/2.:
			// right-side
			switch {
			case ny == h/2.:
				// ignore
			case ny < h/2.:
				// top
				rt++
			case ny > h/2.:
				// bottom
				rb++
			}
		}
	}
	printGrid(grid)
	println(lt, rt, rb, lb, lt*rt*rb*lb)

	maxN := 0
	i := 0
	for i < 15000 {
		n, grid := printGridAfter(robots, w, h, i)
		if n > maxN {
			maxN = n
			print("\033[2J\033[H")
			println(i)
			printGrid(grid)
		}
		//time.Sleep(250 * time.Millisecond)
		i++
	}
}

func printGridAfter(robots []robot, w, h, s int) (int, [][]byte) {
	grid := makeGrid(h, w)
	for _, r := range robots {
		nx, ny := r.move(w, h, s)
		grid[ny][nx]++
	}

	n := 0
	for i := 1; i < h-2; i++ {
		for j := 1; j < w-2; j++ {
			nl := 0
			if grid[i-1][j] > 0 {
				nl++
			}
			if grid[i+1][j] > 0 {
				nl++
			}
			if grid[i][j-1] > 0 {
				nl++
			}
			if grid[i][j+1] > 0 {
				nl++
			}
			if nl >= 2 {
				n++
			}
		}
	}
	return n, grid
}
