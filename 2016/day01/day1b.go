package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction int

func (d Direction) turn(t string) Direction {
	switch t {
	case "L":
		// Avoid negatives by adding 4 first
		return (d + 4 - 1) % 4
	case "R":
		return (d + 1) % 4
	}
	return NORTH
}

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

type Position struct {
	d    Direction
	x, y int
}

func (p *Position) turn(t string) {
	p.d = p.d.turn(t)
}

func (p *Position) move(steps int) {
	switch p.d {
	case NORTH:
		p.y += steps
	case EAST:
		p.x += steps
	case SOUTH:
		p.y -= steps
	case WEST:
		p.x -= steps
	}
}

func (p *Position) String() string {
	return fmt.Sprintf("%dx%d", p.x, p.y)
}

func parse(s string) (string, int) {
	t := s[0]
	steps, err := strconv.Atoi(s[1:])
	if err != nil {
		log.Fatal(err)
	}
	return string(t), steps
}

func locate(seq string) *Position {
	locations := make(map[string]struct{})

	p := &Position{}

	for _, s := range strings.Split(seq, ", ") {
		t, steps := parse(s)
		p.turn(t)
		for i := 0; i < steps; i++ {
			p.move(1)

			if _, ok := locations[p.String()]; ok {
				return p
			}

			locations[p.String()] = struct{}{}
		}
	}
	return p
}

func main() {
	file, err := os.Open("input-b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p := locate(scanner.Text())
		fmt.Printf("Blocks: %d\n", int64(math.Abs(float64(p.x))+math.Abs(float64(p.y))))
	}
}
