// find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l.
// The elves also need a little extra paper for each present: the area of the smallest side.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func max(d []int) int {
	m := d[0]
	for _, a := range d {
		if a > m {
			m = a
		}
	}
	return m
}

func sum(d []int) int {
	s := 0
	for _, a := range d {
		s += a
	}
	return s
}

func atoi(s []string) []int {
	i := make([]int, len(s))
	for a, b := range s {
		i[a], _ = strconv.Atoi(b)
	}
	return i
}

func compute(s string) int {
	d := atoi(strings.Split(s, "x"))

	return 2*(sum(d)-max(d)) + (d[0] * d[1] * d[2])
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	size := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		size += compute(scanner.Text())
	}
	fmt.Printf("size: %d\n", size)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
