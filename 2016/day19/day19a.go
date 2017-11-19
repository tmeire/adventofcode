package main

import "fmt"

type ring []int

func (r ring) next(i int) int {
	if i == len(r) {
		i = 0
	}
	for r[i] == 0 {
		i++
		if i == len(r) {
			i = 0
		}
	}
	return i
}

func main() {
	ELVES := 3004953

	presents := ring(make([]int, ELVES))
	for i := range presents {
		presents[i] = 1
	}

	i := 0
	for presents[i] != ELVES {
		j := presents.next(i + 1)
		presents[i] += presents[j]
		presents[j] = 0

		if i == ELVES {
			break
		}

		i = presents.next(j + 1)
	}
	fmt.Println("Winning elf:", i+1)
}
