package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start, end uint64
}

func parse(s string) Range {
	ss := strings.Split(s, "-")

	a, _ := strconv.ParseUint(ss[0], 10, 64)
	b, _ := strconv.ParseUint(ss[1], 10, 64)

	return Range{a, b}
}

func load() []Range {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	blacklist := []Range{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		blacklist = append(blacklist, parse(scanner.Text()))
	}
	return blacklist
}

type ByStart []Range

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].start < a[j].start }

func cleanup(blacklist []Range, maxIP uint64) (uint64, uint64) {
	// sort the ranges by start index
	sort.Sort(ByStart(blacklist))

	// merge overlapping or adjecent ranges
	merged := []Range{blacklist[0]}
	for i := 1; i < len(blacklist); i++ {
		if blacklist[i].start <= merged[len(merged)-1].end+1 {
			if merged[len(merged)-1].end < blacklist[i].end {
				merged[len(merged)-1].end = blacklist[i].end
			}
			// Else: blacklist[i] is already fully contained within the last merged range
		} else {
			merged = append(merged, blacklist[i])
		}
	}

	// Smallest allowed ip is the first ip after the first blacklisted range
	smallest := merged[0].end + 1

	// Count the number of allowed ip's
	var allowed uint64
	for i := 0; i < len(merged)-1; i++ {
		allowed += (merged[i+1].start - merged[i].end - 1)
	}
	allowed += (maxIP - merged[len(merged)-1].end)

	return smallest, allowed
}

func main() {
	blacklist := []Range{
		{5, 8},
		{0, 2},
		{4, 7},
	}

	smallest, allowed := cleanup(blacklist, 9)
	fmt.Println("Smallest ip:", smallest)
	fmt.Println("Allowed ip's:", allowed)

	blacklist = load()

	smallest, allowed = cleanup(blacklist, 4294967295)
	fmt.Println("Smallest ip:", smallest)
	fmt.Println("Allowed ip's:", allowed)

}
