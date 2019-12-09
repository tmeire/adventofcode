package main

import (
	"bufio"
	"fmt"
	"os"
)

func countChars(s string) map[rune]int {
	cc := make(map[rune]int)

	for _, c := range s {
		cc[c]++
	}
	return cc
}

func hasTwoThree(cc map[rune]int) (bool, bool) {
	hasTwo := false
	hasThree := false
	for _, v := range cc {
		if v == 2 {
			hasTwo = true
		}
		if v == 3 {
			hasThree = true
		}
	}
	return hasTwo, hasThree
}

func partA(ids []string) int {
	twos := 0
	threes := 0

	for _, id := range ids {
		hasTwo, hasThree := hasTwoThree(countChars(id))
		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}
	return twos * threes
}

func diff(a string, b string) int {
	diff := 0
	for i, r := range []byte(a) {
		if b[i] != r {
			diff++
		}
	}
	return diff
}

func match(a, b string) string {
	rs := make([]byte, 0)
	for i, r := range []byte(a) {
		if b[i] == r {
			rs = append(rs, r)
		}
	}
	return string(rs)
}

func partB(vs []string) string {
	for i, v1 := range vs {
		for _, v2 := range vs[i:] {
			d := diff(v1, v2)
			if d == 1 {
				return match(v1, v2)
			}
		}
	}
	return ""
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	values := []string{}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	fmt.Printf("Part A: final %d\n", partA(values))
	fmt.Printf("Part B: %s\n", partB(values))
}
