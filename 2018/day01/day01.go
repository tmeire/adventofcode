package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partA(vs []int) int {
	total := 0
	for _, v := range vs {
		total += v
	}
	return total
}

func partB(vs []int) int {
	freqs := make(map[int]struct{})

	total := 0
	for {
		for _, v := range vs {
			total += v
			if _, ok := freqs[total]; ok {
				return total
			}
			freqs[total] = struct{}{}
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	values := []int{}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		values = append(values, v)
	}

	fmt.Printf("Part A: final %d\n", partA(values))
	fmt.Printf("Part B: first %d frequency\n", partB(values))
}
