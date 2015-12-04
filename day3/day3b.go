package main

import (
	"fmt"
	"io/ioutil"
)

type Santa struct {
	x      int
	y      int
	houses map[string]bool
}

func (s *Santa) move(b byte) {
	switch b {
	case 60: // <
		s.x += 1
	case 62: // >
		s.x -= 1
	case 94: // ^
		s.y += 1
	case 118: // v
		s.y -= 1
	}
	s.houses[fmt.Sprintf("%dx%d", s.x, s.y)] = true
}

func main() {
	bs, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic("No such file")
	}
	houses := make(map[string]bool)
	houses["0x0"] = true
	s1 := Santa{0, 0, houses}
	s2 := Santa{0, 0, houses}
	for i, b := range bs {
		if i%2 == 0 {
			s1.move(b)
		} else {
			s2.move(b)
		}
	}
	fmt.Println("houses: ", len(houses))
}
