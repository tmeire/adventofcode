package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func minmax(s []int) (int, int) {
	min := -1
	max := 0

	for _, v := range s {
		if min == -1 || v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func evendiv(s []int) int {
	sort.Sort(sort.IntSlice(s))

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j]%s[i] == 0 {
				return s[j] / s[i]
			}
		}
	}
	// We should not be able to get here
	return -1
}

func readInts(s string) []int {
	ss := strings.Fields(s)

	ints := []int{}
	for _, f := range ss {
		a, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		ints = append(ints, a)
	}
	return ints
}

func readFile(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Solve() {
	if len(os.Args) < 2 {
		fmt.Println("Expected the input file name as commandline argument.")
		os.Exit(1)
	}

	lines := readFile(os.Args[1])
	ints := [][]int{}
	for _, l := range lines {
		ints = append(ints, readInts(l))
	}

	sum := 0
	for _, i := range ints {
		min, max := minmax(i)
		sum += (max - min)
	}
	fmt.Println("Part A", sum)

	divsum := 0
	for _, i := range ints {
		div := evendiv(i)
		divsum += div
	}
	fmt.Println("Part B", divsum)
}
