package fields

import "math/bits"

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
