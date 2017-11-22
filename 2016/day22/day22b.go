package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/cnf/structhash"
)

type node struct {
	x    int
	y    int
	size int
	used int
	free int

	goal bool
}

func (a *node) move(b *node) {
	//fmt.Println(a, b)
	if a.used != 0 {
		panic("Trying to move to non-empty node.")
	}
	a.used = b.used
	a.free = a.size - a.used

	b.used = 0
	b.free = a.size
	if b.goal {
		a.goal = true
		b.goal = false
	}
}

//   /dev/grid/node-x0-y0     93T   71T    22T   76%
var PATTERN = regexp.MustCompile(`/dev/grid/node-x([0-9]+)-y([0-9]+)\s+([0-9]+)T\s+([0-9]+)T\s+([0-9]+)T\s+([0-9]+)%`)

func parse(l string) node {
	ss := PATTERN.FindStringSubmatch(l)

	x, _ := strconv.Atoi(ss[1])
	y, _ := strconv.Atoi(ss[2])
	s, _ := strconv.Atoi(ss[3])
	u, _ := strconv.Atoi(ss[4])
	f, _ := strconv.Atoi(ss[5])

	n := node{x, y, s, u, f, false}
	return n
}

func load() ([]node, int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nodes := []node{}

	maxX, maxY := 0, 0

	scanner := bufio.NewScanner(file)
	// skip the first to lines
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		n := parse(scanner.Text())
		if n.x > maxX {
			maxX = n.x
		}
		if n.y > maxY {
			maxY = n.y
		}
		nodes = append(nodes, n)
	}
	return nodes, maxX, maxY
}

type grid struct {
	n []node

	maxX  int
	maxY  int
	zeroX int
	zeroY int
	goalX int
	goalY int

	cost int

	next *grid
}

func (g *grid) clone() *grid {
	c := *g
	c.n = append([]node(nil), c.n...)
	return &c
}

func (g grid) idx(x, y int) int {
	return x*(g.maxY+1) + y
}

func (g grid) move(x, y int) *grid {
	if g.zeroX+x < 0 || g.zeroX+x > g.maxX || g.zeroY+y < 0 || g.zeroY+y > g.maxY {
		return nil
	}

	gc := g.clone()
	z := gc.n[g.idx(g.zeroX, g.zeroY)]
	t := gc.n[g.idx(g.zeroX+x, g.zeroY+y)]

	if z.size >= t.used {
		z.move(&t)
		gc.n[g.idx(g.zeroX, g.zeroY)] = z
		gc.n[g.idx(g.zeroX+x, g.zeroY+y)] = t
		//fmt.Println(gc.n[g.idx(g.zeroX+x, g.zeroY+y)])
		gc.zeroX = t.x
		gc.zeroY = t.y
		if z.goal {
			gc.goalX = z.x
			gc.goalY = z.y
		}
		gc.cost++
		return gc
	}
	return nil
}

func (g grid) moves() []*grid {
	m := make([]*grid, 0, 4)
	g1 := g.move(-1, 0)
	if g1 != nil {
		m = append(m, g1)
	}
	g2 := g.move(0, -1)
	if g2 != nil {
		m = append(m, g2)
	}
	g3 := g.move(1, 0)
	if g3 != nil {
		m = append(m, g3)
	}
	g4 := g.move(0, +1)
	if g4 != nil {
		m = append(m, g4)
	}
	return m
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (g *grid) score() int {
	// NOTE: If you're trying to use this on your input,
	// you may want to update the function to match your grid better.

	// get around the barrier of unmovable data, but add a large weight to it
	// to make sure it ends up at the back of the queue when options beyond the
	// wall are added.
	if g.zeroY >= 6 && g.zeroX > 7 {
		return 1000 + g.zeroX
	}

	if g.zeroY > 0 {
		return 50 + g.zeroY
	}

	// move the empty spot right in front of the spot with the goal data
	return (g.maxX - g.zeroX)
}

func (g *grid) add(n *grid) {
	s := n.score()

	c := g
	for c.next != nil && c.next.score() < s {
		c = c.next
	}
	n.next = c.next
	c.next = n
}

func (g *grid) hash() string {
	return fmt.Sprintf("%x", structhash.Md5(g, 1))
}

func (g *grid) print() {
	for j := 0; j <= g.maxY; j++ {
		for i := 0; i <= g.maxX; i++ {
			n := g.n[g.idx(i, j)]
			if n.goal {
				fmt.Print(" G ")
			} else if n.used == 0 {
				fmt.Print(" _ ")
			} else if n.used > 90 {
				fmt.Print(" * ")
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}
	fmt.Println("---------")
}

func main() {
	nodes, maxX, maxY := load()

	idx := func(x, y int) int {
		return x*(maxY+1) + y
	}

	zeroX, zeroY := 0, 0

	nmap := make([]node, (maxX+1)*(maxY+1))
	for _, n := range nodes {
		//fmt.Println(n.x, n.y, idx(n.x, n.y))
		if n.used == 0 {
			zeroX = n.x
			zeroY = n.y
		}
		if n.x == maxX && n.y == 0 {
			n.goal = true
		}
		nmap[idx(n.x, n.y)] = n
	}

	hashes := make(map[string]struct{})

	g := &grid{nmap, maxX, maxY, zeroX, zeroY, 0, maxX, 0, nil}
	hashes[g.hash()] = struct{}{}

	g.print()
	fmt.Println("You may want to count by hand. It's probably faster.")

	i := 0
	for g != nil {
		i++
		for _, m := range g.moves() {
			if m.zeroX == g.goalX-1 && m.zeroY == 0 {
				fmt.Println(m.cost + 1 + (g.goalX-1)*5)
				return
			}

			if _, ok := hashes[m.hash()]; ok {
				continue
			}

			hashes[g.hash()] = struct{}{}
			g.add(m)
		}
		g = g.next
	}

	// mv (maxX, 0) to (0,0)

	// find the shortest path from (zeroX, zeroY) to (maxX-1, 0)
	// move G from (maxX, 0) to (maxX-1,0)
	// (zeroX, zeroY) = (maxX, 0)

	// find the shortest path from (maxX-1, 0)
	// move G from (maxX-1, 0) to (maxX-2,0)
}
