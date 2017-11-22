package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type node struct {
	name string
	size int
	used int
	free int
}

//   /dev/grid/node-x0-y0     93T   71T    22T   76%
var PATTERN = regexp.MustCompile(`([a-z0-9/]+)\s+([0-9]+)T\s+([0-9]+)T\s+([0-9]+)T\s+([0-9]+)%`)

func parse(l string) node {
	ss := PATTERN.FindStringSubmatch(l)

	s, _ := strconv.Atoi(ss[2])
	u, _ := strconv.Atoi(ss[3])
	f, _ := strconv.Atoi(ss[4])

	return node{ss[1], s, u, f}
}

func load() []node {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nodes := []node{}

	scanner := bufio.NewScanner(file)
	// skip the first to lines
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		nodes = append(nodes, parse(scanner.Text()))
	}
	return nodes
}
func main() {
	nodes := load()

	viables := 0
	for i, a := range nodes {
		for j := i + 1; j < len(nodes); j++ {
			if a.used > 0 && a.used <= nodes[j].free {
				viables++
			}
			if nodes[j].used > 0 && nodes[j].used <= a.free {
				viables++
			}
		}
	}
	fmt.Println("Viable couples:", viables)
}
