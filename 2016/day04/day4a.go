package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var REG = regexp.MustCompile(`([a-z\-]+)-([0-9]+)\[([a-z]+)\]`)

func mustInt(a int, err error) int {
	if err != nil {
		panic(err)
	}
	return a
}

type char struct {
	b rune
	c int
}

type charslice []char

func (cs charslice) Len() int {
	return len(cs)
}

func (cs charslice) Less(i, j int) bool {
	if cs[i].c < cs[j].c {
		return false
	} else if cs[i].c > cs[j].c {
		return true
	} else {
		return cs[i].b < cs[j].b
	}
}

func (cs charslice) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func checksum(s string) string {
	data := make(map[rune]int)
	for _, ss := range s {
		if ss != '-' {
			data[ss]++
		}
	}

	counts := make([]char, 0, len(data))
	for k, v := range data {
		counts = append(counts, char{k, v})
	}

	sort.Sort(charslice(counts))

	sum := make([]rune, 5)
	for i, c := range counts[:5] {
		sum[i] = c.b
	}
	return string(sum)
}

func verify(s, chck string) bool {
	return chck == checksum(s)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := REG.FindStringSubmatch(scanner.Text())

		if verify(s[1], s[3]) {
			sum += mustInt(strconv.Atoi(s[2]))
		}
	}
	fmt.Println("Sum of sector ID: ", sum)
}
