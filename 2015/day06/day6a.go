package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x int
	y int
}

type grid struct {
	state [1000][1000]bool
}

func NewGrid() *grid {
	return new(grid)
}

func (g *grid) On(s Point, e Point) {
	for i := s.x; i < e.x+1; i++ {
		for j := s.y; j < e.y+1; j++ {
			g.state[i][j] = true
		}
	}
}

func (g *grid) Toggle(s Point, e Point) {
	for i := s.x; i < e.x+1; i++ {
		for j := s.y; j < e.y+1; j++ {
			g.state[i][j] = !g.state[i][j]
		}
	}
}

func (g *grid) Off(s Point, e Point) {
	for i := s.x; i < e.x+1; i++ {
		for j := s.y; j < e.y+1; j++ {
			g.state[i][j] = false
		}
	}
}

func (g *grid) Count() int {
	c := 0
	for _, i := range g.state {
		for _, j := range i {
			if j {
				c += 1
			}
		}
	}
	return c
}

var regex = regexp.MustCompile("([a-z]*) ([0-9]*),([0-9]*) through ([0-9]*),([0-9]*)")

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func parse(action string) (string, Point, Point) {
	m := regex.FindAllStringSubmatch(action, -1)
	if m == nil || len(m[0]) != 6 {
		return "", Point{}, Point{}
	}
	return m[0][1], Point{atoi(m[0][2]), atoi(m[0][3])}, Point{atoi(m[0][4]), atoi(m[0][5])}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	g := NewGrid()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		action, x, y := parse(scanner.Text())
		switch action {
		case "on":
			g.On(x, y)
		case "toggle":
			g.Toggle(x, y)
		case "off":
			g.Off(x, y)
		}
	}
	fmt.Printf("Lit lights: %d\n", g.Count())
}
