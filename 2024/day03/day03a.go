package day03

import (
	"os"
	"regexp"
	"strconv"
)

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

var mul = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
var all = regexp.MustCompile(`(do)\(\)|(don't)\(\)|(mul)\(([0-9]+),([0-9]+)\)`)
var ops = regexp.MustCompile(`(do|don't|mul)\((([0-9]+),([0-9]+))?\)`)

func partA(prog string) int {
	muls := mul.FindAllStringSubmatch(string(prog), -1)

	sum := 0
	for _, mul := range muls {
		a := atoi(mul[1])
		b := atoi(mul[2])
		sum += a * b
	}
	return sum
}

func partB(prog string) int {
	ins := all.FindAllStringSubmatch(string(prog), -1)

	sum := 0
	enabled := true
	for _, in := range ins {
		switch {
		case in[1] != "": // do
			enabled = true
		case in[2] != "": // don't
			enabled = false
		case in[3] != "": // mul
			if enabled {
				a := atoi(in[4])
				b := atoi(in[5])
				sum += a * b
			}
		}
	}
	return sum
}

func partOps(prog string) int {
	ins := ops.FindAllStringSubmatch(string(prog), -1)

	sum := 0
	enabled := true
	for _, in := range ins {
		switch in[1] {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if enabled {
				a := atoi(in[3])
				b := atoi(in[4])
				sum += a * b
			}
		}
	}
	return sum
}

func Solve() {
	prog, err := os.ReadFile("./2024/day03/example2.txt")
	if err != nil {
		panic(err)
	}

	sum := partA(string(prog))
	println(sum)

	sum = partB(string(prog))
	println(sum)

	sum = partOps(string(prog))
	println(sum)
}
