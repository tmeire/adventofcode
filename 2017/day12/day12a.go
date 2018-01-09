package day12

import (
	"fmt"
	"os"

	"github.com/blackskad/adventofcode/collection"
)

func Solve() {
	if len(os.Args) < 2 {
		fmt.Println("Expected the input file name as commandline argument.")
		os.Exit(1)
	}

	pipes := collection.NewEdgeSetFromFile(os.Args[1])

	sets := pipes.NodeSets()
	for _, set := range sets {
		if set.Contains("0") {
			fmt.Println("Part A:", len(set))
			break
		}
	}
	fmt.Println("Part B:", len(sets))
}
