package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func reduce(polymer []int8) []int8 {
	firstRun := true
	reduced := make([]int8, len(polymer))
	for firstRun || len(polymer) != len(reduced) {
		length := 0
		for i := 0; i < len(polymer); i++ {
			if i == len(polymer)-1 || polymer[i]+polymer[i+1] != 0 {
				reduced[length] = polymer[i]
				length++
			} else {
				i++
			}
		}
		//fmt.Printf("%#v\n", reduced)
		polymer, reduced = reduced[0:length], polymer
		firstRun = false
	}
	return reduced
}

func clean(polymer []int8, unit int8) []int8 {
	cleaned := make([]int8, len(polymer))
	length := 0
	for i := 0; i < len(polymer); i++ {
		if polymer[i] != unit && polymer[i] != -unit {
			cleaned[length] = polymer[i]
			length++
		}
	}
	return cleaned[0:length]
}

func abs(x int8) int8 {
	if x < 0 {
		return -x
	}
	return x
}

func encode(polymer []byte) ([]int8, int8) {
	maxUnit := int8(0)
	intpolymer := make([]int8, len(polymer))
	for i := range polymer {
		if polymer[i] <= 90 {
			intpolymer[i] = 64 - int8(polymer[i])
			if maxUnit < abs(intpolymer[i]) {
				maxUnit = abs(intpolymer[i])
			}
		} else {
			intpolymer[i] = int8(polymer[i]) - 96
			if maxUnit < abs(intpolymer[i]) {
				maxUnit = abs(intpolymer[i])
			}
		}
	}
	return intpolymer, maxUnit
}

func decode(intpolymer []int8) []byte {
	polymer := make([]byte, len(intpolymer))
	for i := range intpolymer {
		if intpolymer[i] < 0 {
			polymer[i] = byte(64 - intpolymer[i])
		} else {
			polymer[i] = byte(intpolymer[i] + 96)
		}
	}
	return polymer
}

// y = 64 - x
// y - 64 = -x
// -y + 64 = x

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	polymer, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	//polymer = []byte("dabAcCaCBAcCcaDA")
	//polymer = []byte("ABCDabcd")

	// Normalize the polymer such that the sum of the
	// uppercase and lowercase characters are 0
	intpolymer, maxUnit := encode(polymer)

	reduced := reduce(intpolymer)
	fmt.Printf("Part A: %d\n", len(reduced))

	minLenght := len(intpolymer)
	for u := int8(1); u <= maxUnit; u++ {
		intpolymer, _ := encode(polymer)

		reduced := reduce(clean(intpolymer, u))
		if len(reduced) < minLenght {
			minLenght = len(reduced)
		}
	}

	fmt.Printf("Part B: %d\n", minLenght)
}

/*
4, 1, 2, -2, -1, 1, -4, -1, 1, -4, -1
4, 1, 2,     -1, 1, -2, -1, 1, -4, -1
dabAaBAaDA

dabCBAcaDAcCcaDA

d  a  b   C   B   A  c  a   D   A  c   C  c  a   D   A
4, 1, 2, -3, -2, -1, 3, 1, -4, -1, 3, -3, 3, 1, -4, -1}
*/
