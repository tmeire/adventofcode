package day11

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	n *Node
	v int
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func blink(l *Node) {
	n := l
	for n != nil {
		if n.v == 0 {
			n.v = 1
		} else {
			vs := fmt.Sprint(n.v)
			if len(vs)%2 == 0 {
				n.v = atoi(vs[:len(vs)/2])
				nn := &Node{
					n: n.n,
					v: atoi(vs[(len(vs) / 2):]),
				}
				n.n = nn
				n = nn
			} else {
				n.v *= 2024
			}
		}
		n = n.n
	}
}

var memo = make(map[string]int)

func blinkRec(v int, depth int) (c int) {
	if depth == 75 {
		return 1
	}

	h := fmt.Sprintf("%d-%d", depth, v)
	if vm, ok := memo[h]; ok {
		return vm
	}

	defer func() {
		memo[h] = c
	}()

	if v == 0 {
		return blinkRec(1, depth+1)
	}

	vs := fmt.Sprint(v)
	if len(vs)%2 == 0 {
		return blinkRec(atoi(vs[:len(vs)/2]), depth+1) + blinkRec(atoi(vs[(len(vs)/2):]), depth+1)
	}

	return blinkRec(v*2024, depth+1)
}

func length(current *Node) int {
	var c int
	for current != nil {
		c++
		current = current.n
	}
	return c
}

func Solve() {
	input := "125 17"
	input = "5688 62084 2 3248809 179 79 0 172169"

	parts := strings.Split(input, " ")
	head := &Node{
		v: atoi(parts[0]),
	}
	current := head
	for i := 1; i < len(parts); i++ {
		current.n = &Node{
			v: atoi(parts[i]),
		}
		current = current.n
	}

	for i := 0; i < 25; i++ {
		blink(head)
		println(i, length(head))
	}

	var c int
	current = head
	for current != nil {
		c++
		current = current.n
	}
	println(c)

	// Part 2, solve it recursively with memoization
	var count int
	for i := 0; i < len(parts); i++ {
		count += blinkRec(atoi(parts[i]), 0)
	}
	println(count)
}
