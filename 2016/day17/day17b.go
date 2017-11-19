package main

import (
	"crypto/md5"
	"fmt"
)

func hash(h string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(h)))
}

func isOpen(b byte) bool {
	return b >= 98
}

type Direction int

const (
	UP    Direction = 0
	DOWN  Direction = 1
	LEFT  Direction = 2
	RIGHT Direction = 3
)

type path struct {
	passcode string
	p        string
	x, y     int

	hash string

	next *path
}

func (p *path) score() int {
	return len(p.p) + (3 - p.x) + (0 - p.y)
}

func (p *path) queue(n *path) {
	score := n.score()

	q := p
	for q.next != nil && q.next.score() < score {
		q = q.next
	}
	n.next = q.next
	q.next = n
}

func (p *path) canMove(d Direction) bool {
	fmt.Println(p.passcode + p.p)
	if p.hash == "" {
		p.hash = hash(p.passcode + p.p)
	}

	fmt.Println(p.hash, isOpen(p.hash[d]))
	return isOpen(p.hash[d])
}

func find(passcode string) string {
	node := &path{passcode, "", 0, 3, "", nil}

	for node != nil {
		fmt.Println(node)
		if node.y < 3 && node.canMove(UP) {
			node.queue(&path{passcode, node.p + "U", node.x, node.y + 1, "", nil})
		}
		if node.y > 0 && node.canMove(DOWN) {
			if node.x == 3 && node.y-1 == 0 {
				return node.p + "D"
			}
			node.queue(&path{passcode, node.p + "D", node.x, node.y - 1, "", nil})
		}
		if node.x > 0 && node.canMove(LEFT) {
			node.queue(&path{passcode, node.p + "L", node.x - 1, node.y, "", nil})
		}
		if node.x < 3 && node.canMove(RIGHT) {
			if node.x+1 == 3 && node.y == 0 {
				return node.p + "R"
			}
			node.queue(&path{passcode, node.p + "R", node.x + 1, node.y, "", nil})
		}
		node = node.next
	}
	return ""
}

func main() {
	tests := map[string]string{
		"ihgpwlah": "DDRRRD",
		"kglvqrro": "DDUDRLRRUDRD",
		"ulqzkmiv": "DRURDRUDDLLDLUURRDULRLDUUDDDRR",
	}

	for pass, path := range tests {
		res := find(pass)
		if res != path {
			fmt.Println(pass, path, res)
		}
	}

	fmt.Println(find("pgflpeqp"))
}
