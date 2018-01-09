package day11

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func distance(x, y int) int {
	return (abs(-x) + abs(-x-y) + abs(-y)) / 2
}

func walk(in []string) (int, int) {
	max := 0

	var d, x, y int
	for _, i := range in {
		switch i {
		case "nw":
			x -= 1
		case "se":
			x += 1

		case "n":
			y -= 1
		case "s":
			y += 1

		case "ne":
			x += 1
			y -= 1
		case "sw":
			x -= 1
			y += 1
		}
		d = distance(x, y)
		if d > max {
			max = d
		}
	}
	return d, max
}

func Solve() {
	if len(os.Args) < 2 {
		fmt.Println("Expected the input file name as commandline argument.")
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	o, max := walk(strings.Split(string(b), ","))
	fmt.Println("Part A", o)
	fmt.Println("Part B", max)
}
