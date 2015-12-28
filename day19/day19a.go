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

func parse(s string) *Replacement {
	var r Replacement
	fmt.Sscanf(s, "%s => %s", &(r.o), &(r.r))
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
	molecule := scanner.Text()

	molecules := make(map[string]bool)
	for _, repl := range replacements {
		i := 0
		a := strings.Index(molecule[i:], repl.o)
		for a != -1 {
			x := molecule[0:i] + strings.Replace(molecule[i:], repl.o, repl.r, 1)
			molecules[x] = true
			i += a + 1
			a = strings.Index(molecule[i:], repl.o)
		}
	}
	fmt.Println(len(molecules))
}
