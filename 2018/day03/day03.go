package main

import (
	"bufio"
	"fmt"
	"os"
)

type fabric struct {
	d map[int]map[int]int
}

func (f fabric) claim(c claim) {
	for i := c.x; i < c.x+c.w; i++ {
		for j := c.y; j < c.y+c.h; j++ {
			f.mark(c.id, i, j)
		}
	}
}

func (f fabric) mark(id, x, y int) {
	v := f.d[x]
	if v == nil {
		v = make(map[int]int)
		f.d[x] = v
	}
	f.d[x][y]++
}

func (f fabric) countDoubleMarked() int {
	doubleMarked := 0
	for _, ms := range f.d {
		for _, m := range ms {
			if m > 1 {
				doubleMarked++
			}
		}
	}
	return doubleMarked
}

func (f fabric) findFreestandingClaim(claims []claim) int {
	for _, c := range claims {
		overlay := false
		for i := c.x; i < c.x+c.w; i++ {
			for j := c.y; j < c.y+c.h; j++ {
				if f.d[i][j] > 1 {
					overlay = true
				}
			}
		}
		if !overlay {
			return c.id
		}
	}
	return -1
}

type claim struct {
	id, x, y, w, h int
}

func parseClaim(s string) claim {
	c := claim{}
	_, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &(c.id), &(c.x), &(c.y), &(c.w), &(c.h))
	if err != nil {
		println(s)
		panic(err)
	}

	return c
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	claims := []claim{}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		claims = append(claims, parseClaim(scanner.Text()))
	}

	fabric := fabric{make(map[int]map[int]int)}
	for _, c := range claims {
		fabric.claim(c)
	}

	fmt.Printf("Double claimed inches: %d\n", fabric.countDoubleMarked())
	fmt.Printf("Freestanding claim: %d\n", fabric.findFreestandingClaim(claims))
}
