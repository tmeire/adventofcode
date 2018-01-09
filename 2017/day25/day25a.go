package day25

import "fmt"

type direction bool

const (
	LEFT  direction = true
	RIGHT           = false
)

type statecase struct {
	writeVal  int
	dir       direction
	nextstate *state
}

type state struct {
	zero, one statecase
}

type node struct {
	value       int
	left, right *node
}

type tape struct {
	head  *node
	state *state
}

func (t *tape) write(val int) {
	// write val to the current position
	t.head.value = val
}

func (t *tape) move(d direction) {
	// move the tapehead one step in the direction d
	if d == LEFT {
		if t.head.left == nil {
			t.head.left = &node{right: t.head}
		}
		t.head = t.head.left
	} else {
		if t.head.right == nil {
			t.head.right = &node{left: t.head}
		}
		t.head = t.head.right
	}
}

func (t *tape) step() {
	var sc statecase
	if t.head.value == 0 {
		sc = t.state.zero
	} else {
		sc = t.state.one
	}

	t.write(sc.writeVal)
	t.move(sc.dir)
	t.state = sc.nextstate
}

func (t *tape) ones() int {
	count := 0

	n := t.head
	for n != nil {
		if n.value == 1 {
			count++
		}
		n = n.left
	}
	n = t.head.right
	for n != nil {
		if n.value == 1 {
			count++
		}
		n = n.right
	}
	return count
}

func Solve() {

	a := &state{zero: statecase{1, RIGHT, nil}, one: statecase{0, LEFT, nil}}
	b := &state{zero: statecase{1, LEFT, nil}, one: statecase{0, RIGHT, nil}}
	c := &state{zero: statecase{1, RIGHT, nil}, one: statecase{0, LEFT, nil}}
	d := &state{zero: statecase{1, LEFT, nil}, one: statecase{1, LEFT, nil}}
	e := &state{zero: statecase{0, RIGHT, nil}, one: statecase{0, RIGHT, nil}}
	f := &state{zero: statecase{1, RIGHT, nil}, one: statecase{1, RIGHT, nil}}

	a.zero.nextstate = b
	a.one.nextstate = b

	b.zero.nextstate = c
	b.one.nextstate = e

	c.zero.nextstate = e
	c.one.nextstate = d

	d.zero.nextstate = a
	d.one.nextstate = a

	e.zero.nextstate = a
	e.one.nextstate = f

	f.zero.nextstate = e
	f.one.nextstate = a

	t := &tape{head: &node{}, state: a}

	steps := 12683008
	for i := 0; i < steps; i++ {
		t.step()
	}
	fmt.Println("Part A", t.ones())
}
