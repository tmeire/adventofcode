package day17

import "fmt"

const INPUT = 349

type RingBuffer struct {
	value int
	next  *RingBuffer
}

func Solve() {
	zero := &RingBuffer{0, nil}

	buf := zero //&RingBuffer{0, nil}
	buf.next = buf

	for i := 1; i < 2018; i++ {
		// step before inserting
		for j := 0; j < INPUT; j++ {
			buf = buf.next
		}

		// insert
		buf.next = &RingBuffer{i, buf.next}
		buf = buf.next
	}
	fmt.Println("Part A:", buf.next.value)

	for i := 2018; i < 50000000; i++ {
		// step before inserting
		for j := 0; j < INPUT; j++ {
			buf = buf.next
		}

		// insert
		buf.next = &RingBuffer{i, buf.next}
		buf = buf.next
	}
	fmt.Println("Part B:", zero.next.value)
}
