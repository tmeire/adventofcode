package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var MAXDEPTH = 10

func atoi(s string) uint16 {
	i, _ := strconv.ParseUint(s, 10, 16)
	return uint16(i)
}

type Definition struct {
	in1 string
	in2 string
	out string
	op  string
}

func NewDefinition(s []string) *Definition {
	return &Definition{s[0], s[2], s[3], s[1]}
}

var vals map[string]uint16

var regex = regexp.MustCompile("([a-z0-9]*?)?[ ]?([A-Z]*)?[ ]?([a-z0-9]+) -> ([a-z]*)")

func main() {
	defs := make([]*Definition, 0)
	vals = make(map[string]uint16)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := regex.FindStringSubmatch(scanner.Text())
		if m == nil {
			continue
		}
		d := NewDefinition(m[1:])
		v1, err := strconv.ParseUint(d.in1, 10, 16)
		if err == nil {
			vals[d.in1] = uint16(v1)
		}
		v2, err := strconv.ParseUint(d.in2, 10, 16)
		if err == nil {
			vals[d.in2] = uint16(v2)
		}
		defs = append(defs, d)
	}

	// For problem b, uncomment this line
	// vals["b"] = 46065

	skips := 0
	for skips != len(defs) {
		skips = 0
		for _, d := range defs {
			if _, ok := vals[d.out]; ok {
				// already computed, skip
				skips += 1
				continue
			}

			switch d.op {
			case "NOT":
				val2, ok2 := vals[d.in2]
				if !ok2 {
					continue
				}
				vals[d.out] = ^val2
			case "AND":
				val1, ok1 := vals[d.in1]
				val2, ok2 := vals[d.in2]
				if !ok1 || !ok2 {
					continue
				}
				vals[d.out] = val1 & val2
			case "OR":
				val1, ok1 := vals[d.in1]
				val2, ok2 := vals[d.in2]
				if !ok1 || !ok2 {
					continue
				}
				vals[d.out] = val1 | val2
			case "LSHIFT":
				val1, ok1 := vals[d.in1]
				val2, ok2 := vals[d.in2]
				if !ok1 || !ok2 {
					continue
				}
				vals[d.out] = val1 << val2
			case "RSHIFT":
				val1, ok1 := vals[d.in1]
				val2, ok2 := vals[d.in2]
				if !ok1 || !ok2 {
					continue
				}
				vals[d.out] = val1 >> val2
			case "":
				val2, ok2 := vals[d.in2]
				if !ok2 {
					continue
				}
				vals[d.out] = val2
			}
		}
	}

	fmt.Printf("Value of a: %d\n", vals["a"])
}
