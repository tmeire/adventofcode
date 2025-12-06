package day06

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func Solve() {
	//lines := read("./2025/day06/input.txt")
	//opi := len(lines) - 1
	//var n int
	//for i := 0; i < len(lines[0]); i++ {
	//	switch lines[opi][i] {
	//	case "+":
	//		m := atoi(lines[0][i])
	//		for j := 1; j < opi; j++ {
	//			m += atoi(lines[j][i])
	//		}
	//		n += m
	//	case "*":
	//		m := atoi(lines[0][i])
	//		for j := 1; j < opi; j++ {
	//			m *= atoi(lines[j][i])
	//		}
	//		n += m
	//	}
	//}
	//println(n)

	sum := func(a, b int) int { return a + b }
	mul := func(a, b int) int { return a * b }

	tlines := readBytes("./2025/day06/input.txt")
	var n int
	var x int
	var op func(int, int) int
	for _, l := range tlines {
		if l.n == -1 {
			n += x
			x = 0
			continue
		}
		switch l.op {
		case '+':
			op = sum
			x = l.n
		case '*':
			op = mul
			x = l.n
		case ' ':
			x = op(x, l.n)
		default:
			panic("Invalid op")
		}
	}
	fmt.Printf("%#v\n", n+x)
}

func read(fn string) [][]string {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var res [][]string

	fb := bufio.NewReader(f)
	for {
		l, _, err := fb.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return res
			}
			panic(err.Error())
		}
		res = append(res, strings.Fields(string(l)))
	}
}

type tline struct {
	n  int
	op byte
}

func readBytes(fn string) []tline {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines [][]byte
	fb := bufio.NewReader(f)
	for {
		l, _, err := fb.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err.Error())
		}

		lines = append(lines, bytes.Clone(l))
	}

	var ln int
	for _, l := range lines {
		if ln < len(l) {
			ln = len(l)
		}
	}

	for i, l := range lines {
		if len(l) < ln {
			lines[i] = append(l, bytes.Repeat([]byte{' '}, ln-len(l))...)
		}
	}

	var res []tline
	for j := 0; j < len(lines[0]); j++ {

		var b []byte
		for i := 0; i < len(lines)-1; i++ {
			b = append(b, lines[i][j])
		}
		b = bytes.TrimSpace(b)

		if len(b) == 0 {
			res = append(res, tline{-1, '.'})
			continue
		}
		res = append(res, tline{atoi(string(b)), lines[len(lines)-1][j]})
	}
	return res
}
