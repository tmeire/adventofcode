package main

import (
	"github.com/tmeire/adventofcode/intcode"
)

var data = []int{
	3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 34, 59, 76, 101, 114, 195, 276, 357, 438, 99999, 3, 9, 1001,
	9, 4, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 101, 2, 9, 9, 102, 4, 9, 9, 1001, 9, 3, 9, 102, 2, 9,
	9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 102, 5, 9, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 1001, 9, 4, 9,
	102, 4, 9, 9, 1001, 9, 4, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9,
	101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102,
	2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2,
	9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9,
	4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3,
	9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9,
	1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002,
	9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9,
	4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4,
	9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3,
	9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9,
	101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9,
	2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99}

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}

	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {

	/*
		maxout := 0
		phases := []int{0, 1, 2, 3, 4}
		Perm(phases, func(phases []int) {
			output := []int{0}
			for _, phase := range phases {
				input := []int{phase, output[0]}
				p := intcode.NewProgram(data)
				output = p.Simulate(input)
			}
			if maxout < output[0] {
				maxout = output[0]
			}
		})
		println(maxout)
	*/
	maxout := 0

	phases := []int{5, 6, 7, 8, 9}
	Perm(phases, func(phases []int) {

		p0 := intcode.NewProgram(data)
		p1 := intcode.NewProgram(data)
		p1.Stdin = p0.Stdout
		p2 := intcode.NewProgram(data)
		p2.Stdin = p1.Stdout
		p3 := intcode.NewProgram(data)
		p3.Stdin = p2.Stdout
		p4 := intcode.NewProgram(data)
		p4.Stdin = p3.Stdout

		go p0.Simulate()
		go p1.Simulate()
		go p2.Simulate()
		go p3.Simulate()
		go p4.Simulate()

		p4.Stdin <- phases[4]
		p3.Stdin <- phases[3]
		p2.Stdin <- phases[2]
		p1.Stdin <- phases[1]
		p0.Stdin <- phases[0]
		p0.Stdin <- 0

		lastout := 0
		done := false
		for !done {
			select {
			case lastout = <-p4.Stdout:
				p0.Stdin <- lastout
			case <-p4.Done:
				done = true
			}
		}

		if maxout < lastout {
			maxout = lastout
		}
	})
	println(maxout)
}
