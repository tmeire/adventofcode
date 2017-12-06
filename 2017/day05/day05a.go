package day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(fname string) []int {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		lines = append(lines, a)
	}
	return lines
}

func step(jumps []int, partb bool) int {
	var steps, ip, oldip int

	for ip >= 0 && ip < len(jumps) {
		oldip = ip

		ip += jumps[ip]

		if partb && jumps[oldip] >= 3 {
			jumps[oldip]--
		} else {
			jumps[oldip]++
		}

		steps++
	}
	return steps
}

func Solve() {
	if len(os.Args) < 2 {
		fmt.Println("Expected the input file name as commandline argument.")
		os.Exit(1)
	}

	ints := readFile(os.Args[1])

	jumps := make([]int, len(ints))

	copy(jumps, ints)
	fmt.Println("Part A", step(jumps, false))

	copy(jumps, ints)
	fmt.Println("Part B", step(jumps, true))
}
