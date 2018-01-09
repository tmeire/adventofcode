package day23

import (
	"fmt"
	"os"

	"github.com/blackskad/adventofcode/algo/cpu"
)

func partB() int {
	nonprimes := 0

	b := (79 * 100) + 100000
	c := b + 17000

	for b < c+1 {
		nonprime := false
		for d := 2; d < b/2; d++ {
			if b%d == 0 {
				nonprime = true
				break
			}
		}
		if nonprime {
			nonprimes++
		}
		b += 17
	}
	return nonprimes
}

func Solve() {
	if len(os.Args) < 2 {
		panic("Must pass input file on commandline.")
	}
	program := cpu.Load(os.Args[1])

	core := cpu.CPU{Debug: true}
	core.Load(program)
	core.Execute()

	fmt.Println("Part A", core.Counts["MULInstruction"])

	fmt.Println("Part B", partB())
}
