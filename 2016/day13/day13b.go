package main

import (
	"fmt"
	"math/bits"
)

const FAVORITE_NUMBER uint64 = 1364

type Field struct {
	x, y, moves uint64

	next *Field
}

func (f *Field) isWall() bool {
	l := uint64(f.x*f.x + 3*f.x + 2*f.x*f.y + f.y + f.y*f.y)

	return bits.OnesCount64(l+FAVORITE_NUMBER)%2 != 0
}

func (f *Field) queue(n *Field) {
	q := f
	for q.next != nil {
		q = q.next
	}
	n.next = q.next
	q.next = n
}

func count() int {
	visited := make([]bool, 100*100)
	visited[101] = true

	locations := 0

	check := func(f *Field) bool {
		idx := f.x*100 + f.y
		if f.moves > 50 || visited[idx] {
			// skip everything that's more than 50 moves away or already visited
			return false
		}
		visited[idx] = true
		return !f.isWall()
	}

	f := &Field{1, 1, 0, nil}
	for f != nil {
		locations += 1
		if f.x > 0 {
			c1 := &Field{f.x - 1, f.y, f.moves + 1, nil}
			if check(c1) {
				f.queue(c1)
			}
		}

		c2 := &Field{f.x + 1, f.y, f.moves + 1, nil}
		if check(c2) {
			f.queue(c2)
		}

		if f.y > 0 {
			c3 := &Field{f.x, f.y - 1, f.moves + 1, nil}
			if check(c3) {
				f.queue(c3)
			}
		}

		c4 := &Field{f.x, f.y + 1, f.moves + 1, nil}
		if check(c4) {
			f.queue(c4)
		}
		f = f.next
	}
	return locations
}

func main() {
	fmt.Println("Reachable places:", count())
}

func abs(a uint64) uint64 {
	if a < 0 {
		return -a
	}
	return a
}
