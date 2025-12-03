package day03

import (
	"fmt"
	"io"
	"os"
)

func rec(bank []int, digits int, j int) int {
	// we've got all the digits we need 🥳
	if digits == 0 {
		return j
	}
	// if there are just enough digits left, just add all of then
	if len(bank) == digits {
		for _, v := range bank {
			j = j*10 + v
		}
		return j
	}
	// find the next largest number until we no longer have enough digits
	maxi := 0
	for i := 0; i <= len(bank)-digits; i++ {
		if bank[i] > bank[maxi] {
			maxi = i
		}
	}
	// recurse on the remaining bank
	return rec(bank[maxi+1:], digits-1, j*10+bank[maxi])
}

func Solve() {
	banks := read("./2025/day03/input.txt")

	var total int
	for _, bank := range banks {
		m := rec(bank, 12, 0)
		total += m
	}
	println(total)
}

func read(fn string) [][]int {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var banks [][]int
	for {
		var bank []int
		var j byte
		for {
			_, err := fmt.Fscanf(f, "%c", &j)
			if err != nil {
				if err == io.EOF {
					return append(banks, bank)
				}
				panic(err.Error())
			}
			if j == '\n' {
				break
			}
			bank = append(bank, int(j-'0'))
		}
		banks = append(banks, bank)
	}
}
