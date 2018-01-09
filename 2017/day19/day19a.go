package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction int

func (d Direction) Opposites() (Direction, Direction) {
	if d == UP || d == DOWN {
		return LEFT, RIGHT
	} else {
		return UP, DOWN
	}
}

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Tubes [][]byte

func (t Tubes) move(i, j int, dir Direction) (int, int) {
	switch dir {
	case UP:
		return i - 1, j
	case DOWN:
		return i + 1, j
	case LEFT:
		return i, j - 1
	case RIGHT:
		return i, j + 1
	}
	return -1, -1
}

func (t Tubes) next(i, j int, dir Direction) (int, int, byte) {
	x, y := t.move(i, j, dir)
	if x < 0 || y < 0 || x >= len(t) || y >= len(t[0]) {
		panic(fmt.Sprintf("YIKES: %d,%d", x, y))
	}

	return x, y, t[x][y]
}

func (t Tubes) Follow(i, j int, dir Direction) ([]byte, int) {
	b := []byte{}

	steps := 0

	x, y, c := t.next(i, j, dir)
	for c != ' ' {
		steps++
		if c == '+' {
			// Switch to another direction
			d1, d2 := dir.Opposites()
			x1, y1 := t.move(x, y, d1)
			if t[x1][y1] != ' ' {
				dir = d1
			} else {
				dir = d2
			}
		} else if c != '-' && c != '|' {
			// Track the letter, no need to change direction
			b = append(b, c)
		}

		i, j = x, y
		x, y, c = t.next(i, j, dir)
	}
	return b, steps + 1
}

func load(fname string) [][]byte {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tubes := [][]byte{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bo := scanner.Bytes()
		bc := make([]byte, len(bo))
		copy(bc, bo)
		tubes = append(tubes, bc)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return tubes
}

func Solve() {
	if len(os.Args) < 2 {
		panic("Must pass input filename as commandline argument.")
	}

	t := Tubes(load(os.Args[1]))
	// find the entrypoint
	var x, y int
	for i, v := range t[0] {
		if v != ' ' {
			x, y = 0, i
		}
	}
	b, steps := t.Follow(x, y, DOWN)
	fmt.Println("Part A", string(b))
	fmt.Println("Part B", steps)
}
