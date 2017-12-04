package day04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type set map[string]struct{}

func (s set) contains(x string) bool {
	_, ok := s[x]
	return ok
}

func (s set) put(x string) {
	s[x] = struct{}{}
}

func bytesort(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func isAnagram(f, s string) bool {
	if len(f) != len(s) {
		return false
	}
	return bytesort(f) == bytesort(s)
}

func validate(p string, strict bool) bool {
	s := make(set)
	for _, f := range strings.Fields(p) {
		if s.contains(f) {
			return false
		}

		if strict {
			for k := range s {
				if isAnagram(f, k) {
					return false
				}
			}
		}

		s.put(f)
	}
	return true
}

func readFile(fname string) []string {
	lines := []string{}

	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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

	valids := 0
	for _, l := range lines {
		if validate(l, false) {
			valids++
		}
	}
	fmt.Println("Part A", valids)

	valids = 0
	for _, l := range lines {
		if validate(l, true) {
			valids++
		}
	}
	fmt.Println("Part B", valids)
}
