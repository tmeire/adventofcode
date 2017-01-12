package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type position struct {
	x, y int
}

func cap(x int) int {
	if x < 0 {
		return 0
	}
	if x > 2 {
		return 2
	}
	return x
}

func (p *position) move(d string) {
	switch d {
	case "U":
		p.y = cap(p.y - 1)
	case "D":
		p.y = cap(p.y + 1)
	case "L":
		p.x = cap(p.x - 1)
	case "R":
		p.x = cap(p.x + 1)
	}
}

func (p *position) asNumber() int {
	return 3*p.y + p.x + 1
}

func main() {
	file, err := os.Open("input-a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p := &position{1, 1}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, s := range scanner.Text() {
			p.move(string(s))
		}
		fmt.Printf("%d", p.asNumber())
	}
	fmt.Println()
}
