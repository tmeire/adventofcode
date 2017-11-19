package main

import "fmt"

type ring []bool

func (r ring) next(i int, c int) int {
	for c > 0 {
		i += 1
		if i == len(r) {
			i = 0
		}
		if r[i] {
			c--
		}
	}
	return i
}

func main() {
	//ELVES := 5
	//ELVES := 9
	ELVES := 3004953

	presents := ring(make([]bool, ELVES))
	for i := range presents {
		presents[i] = true
	}

	j := ELVES / 2
	alt := (ELVES + 1) % 2
	for c := 0; c < ELVES-1; c++ {
		// Remove the unlucky elf
		presents[j] = false

		alt = (alt + 1) % 2

		// Move to the next opponent
		j = presents.next(j, 1+alt)
	}
	fmt.Println("Winning elf:", j+1)
}
