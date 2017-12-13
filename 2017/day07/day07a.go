package day07

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Program struct {
	name     string     // the name, read from input
	weight   int        // the weight, read from input
	ids      []string   // the ids of child programs, read from input
	children []*Program // references to the other child programs
	root     bool       // indicator to track the root node
}

func (n *Program) Weight() (int, int) {
	// Collect the weights of all the child nodes
	weights := []int{}
	for _, c := range n.children {
		w, diff := c.Weight()
		if diff != 0 {
			return 0, diff
		}
		weights = append(weights, w)
	}

	sum := 0
	for idx, i := range weights {
		// Check if this weight matches the other weight
		samevalues := 0
		for _, j := range weights {
			if i == j {
				samevalues++
			}
		}
		// Only matches itself, this is the bad one
		if samevalues == 1 {
			return 0, n.children[idx].weight + weights[(idx+1)%2] - i
		}
		// Ok, matches, add to the sum
		sum += i
	}

	return sum + n.weight, 0
}

func parse(s string) *Program {
	p := Program{root: true}

	ss := strings.Split(s, " -> ")

	_, err := fmt.Sscanf(ss[0], "%s (%d)", &p.name, &p.weight)
	if err != nil {
		panic(err)
	}
	if len(ss) > 1 {
		p.ids = strings.Split(ss[1], ", ")
	}

	return &p
}

func readFile(fname string) map[string]*Program {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := map[string]*Program{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p := parse(scanner.Text())
		lines[p.name] = p
	}
	return lines
}

func Solve() {
	t0 := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("Expected the input file name as commandline argument.")
		os.Exit(1)
	}

	programs := readFile(os.Args[1])
	t := time.Now()
	// Part A
	// Link them all together
	for _, p := range programs {
		for _, c := range p.ids {
			programs[c].root = false
			p.children = append(p.children, programs[c])
		}
	}

	var root *Program
	for _, p := range programs {
		if p.root {
			root = p
			break
		}
	}

	fmt.Println("Part A", root.name)

	// Part B, Find the bad node and its weight diff
	_, diff := root.Weight()

	fmt.Println("Part B", diff)
	fmt.Println(time.Since(t0))
	fmt.Println(time.Since(t))
}
