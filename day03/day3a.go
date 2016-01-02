package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic("No such file")
	}
	houses := make(map[string]bool)
	houses["0x0"] = true
	x := 0
	y := 0
	for _, b := range bs {
		switch b {
		case 60: // <
			x += 1
		case 62: // >
			x -= 1
		case 94: // ^
			y += 1
		case 118: // v
			y -= 1
		}
		houses[fmt.Sprintf("%dx%d", x, y)] = true
	}
	fmt.Println("houses: ", len(houses))
}
