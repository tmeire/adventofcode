package main

import "fmt"

var data = []int{
	1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2,
	1, 9, 19, 1, 13, 19, 23, 2, 23, 9, 27, 1, 6, 27,
	31, 2, 10, 31, 35, 1, 6, 35, 39, 2, 9, 39, 43, 1,
	5, 43, 47, 2, 47, 13, 51, 2, 51, 10, 55, 1, 55, 5,
	59, 1, 59, 9, 63, 1, 63, 9, 67, 2, 6, 67, 71, 1, 5,
	71, 75, 1, 75, 6, 79, 1, 6, 79, 83, 1, 83, 9, 87, 2,
	87, 10, 91, 2, 91, 10, 95, 1, 95, 5, 99, 1, 99, 13,
	103, 2, 103, 9, 107, 1, 6, 107, 111, 1, 111, 5, 115,
	1, 115, 2, 119, 1, 5, 119, 0, 99, 2, 0, 14, 0,
}

const (
	opcodeADD = 1
	opcodeMUL = 2
	opcodeEND = 99
)

func process(data []int, noun, verb int) int {
	data[1] = noun
	data[2] = verb

	for idx := 0; idx < len(data); idx += 4 {
		switch data[idx] {
		case opcodeADD:
			a1 := data[data[idx+1]]
			a2 := data[data[idx+2]]
			data[data[idx+3]] = a1 + a2
		case opcodeMUL:
			a1 := data[data[idx+1]]
			a2 := data[data[idx+2]]
			data[data[idx+3]] = a1 * a2
		case opcodeEND:
			break
		}
	}
	return data[0]
}

func main() {
	d1 := append(make([]int, 0, len(data)), data...)
	println(process(d1, 12, 2))

	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			dcopy := append(make([]int, 0, len(data)), data...)
			res := process(dcopy, a, b)
			if res == 19690720 {
				fmt.Printf("%d\n", 100*a+b)
				return
			}
		}
	}
}
