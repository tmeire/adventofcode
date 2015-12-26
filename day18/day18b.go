package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type Grid struct {
	tmp  [100][100]bool
	curr [100][100]bool
}

func (g *Grid) Load(scanner *bufio.Scanner) {
	for i := 0; i < 100 && scanner.Scan(); i += 1 {
		for j, x := range scanner.Text() {
			g.curr[i][j] = x == '#'
		}
	}
	// set all corners to on
	g.curr[0][0] = true
	g.curr[0][99] = true
	g.curr[99][0] = true
	g.curr[99][99] = true
}

func (g *Grid) backup() {
	for i := 0; i < 100; i += 1 {
		for j := 0; j < 100; j += 1 {
			g.tmp[i][j] = g.curr[i][j]
		}
	}
}

func (g *Grid) check(a int, b int) bool {
	// addition for part b: corners always on
	if (a == 0 && b == 0) || (a == 0 && b == 99) || (a == 99 && b == 0) || (a == 99 && b == 99) {
		return true
	}

	on := 0
	for i := -1; i < 2; i += 1 {
		for j := -1; j < 2; j += 1 {
			// skip the square itself
			if i == 0 && j == 0 {
				continue
			}
			if a+i >= 0 && a+i < 100 && b+j >= 0 && b+j < 100 && g.tmp[a+i][b+j] {
				on += 1
			}
		}
	}
	if g.tmp[a][b] {
		return on == 2 || on == 3
	} else {
		return on == 3
	}
}

func (g *Grid) Animate() {
	g.backup()

	for i := 0; i < 100; i += 1 {
		for j := 0; j < 100; j += 1 {
			g.curr[i][j] = g.check(i, j)
		}
	}
}

func (g *Grid) Count() int {
	c := 0
	for i := 0; i < 100; i += 1 {
		for j := 0; j < 100; j += 1 {
			if g.curr[i][j] {
				c += 1
			}
		}
	}
	return c
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := &Grid{}
	grid.Load(scanner)
	for i := 0; i < 100; i += 1 {
		grid.Animate()
	}
	fmt.Println(grid.Count())
}
