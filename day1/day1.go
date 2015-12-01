package main

import (
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		panic("No such file")
	}
	floor := 0
	for _, b := range bs {
		switch b {
		case 40: // (
			floor += 1
		case 41: // )
			floor -= 1
		}
	}
	print("Floor: ", floor)
}
