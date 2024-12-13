package day13

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	x, y int64
}

type machine struct {
	p, a, b pair
}

func cost(m machine) int64 {
	println(m.p.x)
	var cost int64 = math.MaxInt64
	for i := int64(0); i < 100; i++ {
		for j := int64(0); j < 100; j++ {
			if m.p.x == (i*m.a.x)+(j*m.b.x) && m.p.y == (i*m.a.y)+(j*m.b.y) {
				c := 3*i + j
				if c < cost {
					cost = c
				}
			}
		}
	}
	return cost
}

func cost2(m machine) int64 {
	px := m.p.x + 10000000000000
	py := m.p.y + 10000000000000

	a := ((py * m.b.x) - (px * m.b.y)) / ((m.b.x * m.a.y) - (m.a.x * m.b.y))
	b := (px - (a * m.a.x)) / (m.b.x)

	pxn := (a * m.a.x) + (b * m.b.x)
	pyn := (a * m.a.y) + (b * m.b.y)

	// double check,int64 conversion might have chopped off some bits
	if px == pxn && py == pyn {
		return 3*a + b
	} else {
		println(px, py, pxn, pyn, a, b)
	}
	return math.MaxInt64
}

func atoi(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func readButton(s string) pair {
	return pair{
		x: atoi(s[12:14]),
		y: atoi(s[18:20]),
	}
}

func readPrize(s string) pair {
	p := strings.Split(s, ", ")
	p0 := strings.Split(p[0], "=")
	p1 := strings.Split(p[1], "=")
	return pair{
		x: atoi(p0[1]),
		y: atoi(p1[1]),
	}
}

func read(fileName string) []machine {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	println(len(lines))

	var machines []machine
	for i := 0; i < len(lines)-3; i += 4 {
		machines = append(machines, machine{
			a: readButton(lines[i]),
			b: readButton(lines[i+1]),
			p: readPrize(lines[i+2]),
		})
	}
	return machines
}

func Solve() {
	// a * x + b * x = X
	// a * y + b * y = Y

	machines := read("./2024/day13/input.txt")

	var tokens, tokens2 int64
	for _, machine := range machines {
		c := cost(machine)
		if c != math.MaxInt {
			tokens += c
		}
		c2 := cost2(machine)
		if c2 != math.MaxInt {
			tokens2 += c2
		}
	}
	println(tokens)
	println(tokens2)
}
