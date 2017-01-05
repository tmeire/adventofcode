package main

import "fmt"

func increment(in []byte) []byte {
	i := len(in) - 1
	for i > -1 {
		if in[i] == 'z' {
			in[i] = 'a'
			i -= 1
		} else {
			in[i] += 1
			i = -1
		}
	}
	return in
}

func containsBadChar(in []byte) bool {
	for _, b := range in {
		if b == 'l' || b == 'i' || b == 'o' {
			return true
		}
	}
	return false
}

func containsStraight(in []byte) bool {
	for i := 2; i < len(in); i += 1 {
		if in[i] == in[i-1]+1 && in[i] == in[i-2]+2 {
			return true
		}
	}
	return false
}

func containsPairs(in []byte) bool {
	pairs := 0
	for i := 1; i < len(in); i += 1 {
		if in[i] == in[i-1] {
			pairs += 1
			i += 1
		}
	}
	return pairs >= 2
}

func check(in []byte) bool {
	return !containsBadChar(in) && containsStraight(in) && containsPairs(in)
}

func main() {
	input := increment([]byte("hepxcrrq"))
	for !check(input) {
		input = increment(input)
	}
	fmt.Println(string(input))
	// part b, next valid password
	input = increment(input)
	for !check(input) {
		input = increment(input)
	}
	fmt.Println(string(input))
}
