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

func (f *Field) score() uint64 {
	return f.moves + abs(TARGET_X-f.x) + abs(TARGET_Y-f.y) // TODO Remaining distance estimate ()
}

func (f *Field) queue(n *Field) {
	score := n.score()

	q := f
	for q.next != nil && q.next.score() < score {
		q = q.next
	}
	n.next = q.next
	q.next = n
}

const (
	TARGET_X = 31
	TARGET_Y = 39
)

func search() uint64 {
	visited := make([]bool, 100*100)
	visited[101] = true

	check := func(f *Field) bool {
		idx := f.x*100 + f.y
		if visited[idx] {
			return false
		}
		visited[idx] = true
		return !f.isWall()
	}

	f := &Field{1, 1, 0, nil}
	for f != nil {
		if f.x > 0 {
			c1 := &Field{f.x - 1, f.y, f.moves + 1, nil}
			if c1.x == TARGET_X && c1.y == TARGET_Y {
				return c1.moves
			}
			if check(c1) {
				f.queue(c1)
			}
		}

		c2 := &Field{f.x + 1, f.y, f.moves + 1, nil}
		if c2.x == TARGET_X && c2.y == TARGET_Y {
			return c2.moves
		}
		if check(c2) {
			f.queue(c2)
		}

		if f.y > 0 {
			c3 := &Field{f.x, f.y - 1, f.moves + 1, nil}
			if c3.x == TARGET_X && c3.y == TARGET_Y {
				return c3.moves
			}
			if check(c3) {
				f.queue(c3)
			}
		}

		c4 := &Field{f.x, f.y + 1, f.moves + 1, nil}
		if c4.x == TARGET_X && c4.y == TARGET_Y {
			return c4.moves
		}
		if check(c4) {
			f.queue(c4)
		}
		f = f.next
	}
	return 0
}

func main() {
	fmt.Println("Required moves:", search())
}

func abs(a uint64) uint64 {
	if a < 0 {
		return -a
	}
	return a
}
