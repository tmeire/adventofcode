package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var spec = make(map[string]int)

func init() {
	spec["children"] = 3
	spec["cats"] = 7
	spec["samoyeds"] = 2
	spec["pomeranians"] = 3
	spec["akitas"] = 0
	spec["vizslas"] = 0
	spec["goldfish"] = 5
	spec["trees"] = 3
	spec["cars"] = 2
	spec["perfumes"] = 1
}

type Aunt struct {
	data map[string]int
}

func (a *Aunt) score() int {
	total := 0
	for k, v := range a.data {
		total = sum(total, match(spec[k], v))
	}
	return total
}

func sum(total int, score int) int {
	if total == -1 || score == -1 {
		return -1
	}
	return total + score
}

func match(a int, b int) int {
	if a == -1 {
		return 0
	}
	if a == b {
		return 1
	}
	return -1
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func split(s string) map[string]int {
	m := make(map[string]int)

	d := strings.Split(s, ", ")
	for _, p := range d {
		kv := strings.Split(p, ": ")
		m[strings.TrimSpace(kv[0])] = atoi(kv[1])
	}
	return m
}

func parse(s string) *Aunt {
	return &Aunt{split(s[strings.Index(s, ":")+1:])}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	aunts := make([]*Aunt, 0, 4)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		aunts = append(aunts, parse(scanner.Text()))
	}

	best := 0
	idx := 0
	for i, aunt := range aunts {
		if score := aunt.score(); score > best {
			best = score
			idx = i
		}
	}
	fmt.Println(idx + 1)
}
