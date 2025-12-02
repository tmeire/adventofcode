package day01

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func Solve() {
	steps := read()

	dial := 50
	c := 0
	for _, s := range steps {
		sn := s.n
		switch s.d {
		case 'L':
			for sn > 100 {
				c++
				sn -= 100
			}
			d1 := dial - sn
			if d1 < 0 && dial != 0 {
				c++
			}
			dial = (d1 + 100) % 100
		case 'R':
			for sn > 100 {
				c += 1
				sn -= 100
			}
			d1 := dial + sn
			if d1 > 100 && dial != 0 {
				c++
			}
			dial = d1 % 100
		}
		if dial == 0 {
			c++
		}
	}
	println(c)
}

type step struct {
	d byte
	n int
}

func read() []step {
	f, err := os.Open("./2025/day01/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var steps []step
	for {
		var d byte
		var n int
		_, err := fmt.Fscanf(f, "%c%d\n", &d, &n)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return steps
			}
			panic(err.Error())
		}
		steps = append(steps, step{d, n})
	}
}
