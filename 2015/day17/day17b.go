package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var containers = make([]int, 0, 4372)

func count(sizes []int, start int, total int, items int) int {
	// end of the list, no more combinations
	if start == len(sizes) {
		return 0
	}

	// skip this entry and count all combinations without this one
	c := count(sizes, start+1, total, items)

	// now try all combinations with this item included
	current := total + sizes[start]
	if current == 150 {
		// exactly made the 150 mark, add match, no need to proceed with this one included
		c += 1
		containers = append(containers, items+1)
	} else if total+sizes[start] > 150 {
		// crossed the 150 mark, no match, no need to proceed with this one included
		c += 0
	} else {
		// haven't reached 150 yet, try with the next item
		c += count(sizes, start+1, total+sizes[start], items+1)
	}
	return c
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sizes := make([]int, 0, 20)

	// read all sized in increasing order
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sizes = append(sizes, atoi(scanner.Text()))
	}

	// sort the sizes in increasing order
	sort.Ints(sizes)

	fmt.Println("Number of combo's:", count(sizes, 0, 0, 0))

	c := 0
	min := 20
	for _, x := range containers {
		if x == min {
			c += 1
		} else if x < min {
			min = x
			c = 1
		}
	}
	fmt.Println("Combo's with minimum number of containers:", c)
}
