package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tmeire/adventofcode/io"
)

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func parse(equation string) (int, []int) {
	parts := strings.Split(equation, ": ")

	res := atoi(parts[0])

	var ops []int
	for _, op := range strings.Split(parts[1], " ") {
		ops = append(ops, atoi(op))
	}
	return res, ops
}

func conc(a, b int) int {
	v := atoi(fmt.Sprintf("%d%d", a, b))
	return v
}

func calc(ops []int, res int, req int) bool {
	if len(ops) == 0 {
		return res == req
	}

	a := calc(ops[1:], res+ops[0], req)
	b := calc(ops[1:], res*ops[0], req)

	return a || b
}

func calc_3ops(ops []int, res int, req int) bool {
	if len(ops) == 0 {
		return res == req
	}

	a := calc_3ops(ops[1:], res+ops[0], req)
	b := calc_3ops(ops[1:], res*ops[0], req)
	c := calc_3ops(ops[1:], conc(res, ops[0]), req)

	return a || b || c
}

func Solve() {
	equations, err := io.ReadLinesFromFile("./2024/day07/input.txt")
	if err != nil {
		panic(err)
	}

	var sum, sum3ops int
	for _, eq := range equations {
		res, operands := parse(eq)

		if calc(operands[1:], operands[0], res) {
			sum += res
		}
		if calc_3ops(operands[1:], operands[0], res) {
			sum3ops += res
		}
	}
	println(sum)
	println(sum3ops)
}
