package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type graph map[byte]node

func (g graph) walk() int {
	// Naively walk each arch in the graph
	return g.recWalk([]byte{'0'}, 0)
}

func (g graph) recWalk(path []byte, cost int) int {
	// We've seen all nodes
	if len(path) == len(g) {
		//return cost
		// Part b, add the cost to travel back from the last node to 0
		return cost + g[path[len(path)-1]]['0']
	}

	min := -1
	for b, d := range g[path[len(path)-1]] {
		fmt.Println(path, b)
		if !contains(path, b) {
			c := g.recWalk(append(path, b), cost+d)
			if min == -1 || c < min {
				min = c
			}
		}
	}
	return min
}

type node map[byte]int

func contains(bs []byte, n byte) bool {
	for _, b := range bs {
		if b == n {
			return true
		}
	}
	return false
}

type layout struct {
	f []byte
	w int
	h int

	pos map[byte]int
}

func (l *layout) moves() []int {
	return []int{-l.w, 1, l.w, -1}
}

func (l *layout) graph() *graph {
	g := make(graph)
	for bo, _ := range l.pos {
		g[bo] = make(node)
	}

	for bo, i := range l.pos {
		for bt, j := range l.pos {
			if d, ok := g[bt][bo]; ok {
				g[bo][bt] = d
			} else {
				dist := travel(l, i, j)
				if dist > 0 {
					g[bo][bt] = dist
				}
			}
		}
	}

	return &g
}

type queue struct {
	pos   int
	moves int
	score int

	next *queue
}

func (q *queue) add(n *queue) {
	c := q
	for c.next != nil && c.next.score < n.score {
		c = c.next
	}
	n.next = c.next
	c.next = n
}

func dist(l *layout, a, b int) int {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d/l.w + d%l.w
}

func travel(l *layout, a, b int) int {
	if a == b {
		return 0
	}

	visited := make(map[int]bool)

	q := &queue{a, 0, dist(l, a, b), nil}

	for q != nil {
		for _, i := range l.moves() {
			p := q.pos - i
			if p == b {
				return q.moves + 1
			}
			if l.f[p] != '#' && !visited[p] {
				q.add(&queue{p, q.moves + 1, q.score + dist(l, a, b), nil})
				visited[p] = true
			}
		}

		q = q.next
	}
	return -1
}

func load() *layout {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	g := layout{make([]byte, 0), 0, 0, make(map[byte]int)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bsub := []byte(scanner.Text())

		g.w = len(bsub)

		// find the positions of the numbers
		for i, j := range bsub {
			if j != '#' && j != '.' {
				g.pos[j] = g.w*g.h + i
			}
		}

		g.f = append(g.f, bsub...)
		g.h++
	}
	return &g
}

func main() {
	fmt.Println(load().graph().walk())
}
