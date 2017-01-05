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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
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

	x := d[0] * d[1]
	y := d[1] * d[2]
	z := d[0] * d[2]

	return 2*(x+y+z) + min(x, min(y, z))
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
