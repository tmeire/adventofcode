package main

import "fmt"

type node struct {
	value      int
	next, prev *node
}

type circle struct {
	root, current *node
}

func (c *circle) add(i int) (v int) {
	c.current, v = c.current.add(i)
	//c.print()
	return v
}

func (c *circle) print() {
	fmt.Printf("%2d", c.root.value)
	for cn := c.root.next; cn != c.root; cn = cn.next {
		fmt.Printf(" %2d ", cn.value)
	}
	fmt.Println()
}

func newCircle() *circle {
	c := &node{0, nil, nil}
	c.next = c
	c.prev = c

	return &circle{c, c}
}

func (n *node) add(i int) (*node, int) {
	if i%23 != 0 {
		nn := &node{i, n.next.next, n.next}
		nn.next.prev = nn
		nn.prev.next = nn
		return nn, 0
	}

	removed := n.prev.prev.prev.prev.prev.prev.prev

	removed.next.prev = removed.prev
	removed.prev.next = removed.next

	return removed.next, i + removed.value
}

func winner(players, marbles int) int {
	scores := map[int]int{}

	c := newCircle()

	player := 0
	for i := 1; i <= marbles; i++ {
		scores[player] += c.add(i)
		player = (player + 1) % players
	}

	max := 0
	for _, score := range scores {
		if max < score {
			max = score
		}
	}
	return max
}

func main() {
	/*
		10 players; last marble is worth 1618 points: high score is 8317
		13 players; last marble is worth 7999 points: high score is 146373
		17 players; last marble is worth 1104 points: high score is 2764
		21 players; last marble is worth 6111 points: high score is 54718
		30 players; last marble is worth 5807 points: high score is 37305
	*/
	println(winner(9, 25))
	println(winner(10, 1618))
	println(winner(13, 7999))
	println(winner(17, 1104))
	println(winner(21, 6111))
	println(winner(30, 5807))
	println(winner(412, 71646))
	println(winner(412, 7164600))

	/*
						0 16   8  17   4  18  19   2  20  10  21   5  22  11   1  12   6  13   3  14   7  15
		 0 16   8  17   4  18  19   2  24  20  25  10  21   5  22  11   1  12   6  13   3  14   7  15
		 0 16   8  17   4  18  19   2  24  20(25)  10  21   5 22 11  1 12  6 13  3 14  7 15
	*/
}
