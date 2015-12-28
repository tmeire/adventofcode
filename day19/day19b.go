package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Replacement struct {
	o string
	r string
}

func generate(m string, replacements []*Replacement, steps int) int {
	if m == "e" {
		return steps
	}
	for _, repl := range replacements {
		i := 0
		a := strings.Index(m[i:], repl.o)
		for a != -1 {
			x := m[0:i] + strings.Replace(m[i:], repl.o, repl.r, 1)

			result := generate(x, replacements, steps+1)
			if result != -1 {
				return result
			}

			i += a + 1
			a = strings.Index(m[i:], repl.o)
		}
	}
	return -1
}

func parse(s string) *Replacement {
	var r Replacement
	fmt.Sscanf(s, "%s => %s", &(r.r), &(r.o))
	return &r
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	replacements := make([]*Replacement, 0, 43)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 {
			break
		}
		replacements = append(replacements, parse(t))
	}

	scanner.Scan()
	source := scanner.Text()

	fmt.Println(generate(source, replacements, 0))
}
