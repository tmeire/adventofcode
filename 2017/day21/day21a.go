package day21

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type pixels [][]byte

func (p pixels) print() {
	for _, x := range p {
		for _, y := range x {
			fmt.Printf("%s", string(y))
		}
		fmt.Println()
	}
}

func (p pixels) fill(s string) {
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p); j++ {
			p[i][j] = byte(s[i*(len(p)+1)+j])
		}
	}
}

func (p pixels) OnCount() int {
	on := 0
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p); j++ {
			if p[i][j] == '#' {
				on++
			}
		}
	}
	return on
}

func (p pixels) flipH() string {
	g := len(p)
	sig := bytes.Buffer{}
	for i := g - 1; i >= 0; i-- {
		for j := 0; j < g; j++ {
			sig.WriteByte(p[i][j])
		}
		if i != 0 {
			sig.WriteByte('/')
		}
	}
	return sig.String()
}

func (p pixels) flipV() string {
	g := len(p)
	sig := bytes.Buffer{}
	for i := 0; i < g; i++ {
		for j := g - 1; j >= 0; j-- {
			sig.WriteByte(p[i][j])
		}
		if i != g-1 {
			sig.WriteByte('/')
		}
	}
	return sig.String()
}

func (p pixels) rotate90() string {
	g := len(p)
	sig := bytes.Buffer{}
	for j := 0; j < g; j++ {
		for i := g - 1; i >= 0; i-- {
			sig.WriteByte(p[i][j])
		}
		if j != g-1 {
			sig.WriteByte('/')
		}
	}
	return sig.String()
}

func (p pixels) rotate180() string {
	g := len(p)
	sig := bytes.Buffer{}
	for i := g - 1; i >= 0; i-- {
		for j := g - 1; j >= 0; j-- {
			sig.WriteByte(p[i][j])
		}
		if i != 0 {
			sig.WriteByte('/')
		}
	}
	return sig.String()
}

func (p pixels) rotate270() string {
	g := len(p)
	sig := bytes.Buffer{}
	for j := g - 1; j >= 0; j-- {
		for i := 0; i < g; i++ {
			sig.WriteByte(p[i][j])
		}
		if j != 0 {
			sig.WriteByte('/')
		}
	}
	return sig.String()
}

func expandRotations(newbook map[string]string, k string, v string, size int) {
	var p pixels = makeGrid(size)
	p.fill(k)

	// rotate 90
	newbook[p.rotate90()] = v
	// rotate 180
	newbook[p.rotate180()] = v
	// rotate 270
	newbook[p.rotate270()] = v
}

func expand(rulebook map[string]string) map[string]string {
	newbook := make(map[string]string, len(rulebook))
	for k, v := range rulebook {
		var p pixels
		if len(k) == 5 {
			p = makeGrid(2)
		} else {
			p = makeGrid(3)
		}
		p.fill(k)

		// Add the original back in the new rulebook & add rotations
		newbook[k] = v
		expandRotations(newbook, k, v, len(p))

		// Flip horizontally & add rotations
		flipH := p.flipH()
		newbook[flipH] = v
		expandRotations(newbook, flipH, v, len(p))
		// Flip vertically & add rotations
		flipV := p.flipV()
		newbook[flipV] = v
		expandRotations(newbook, flipV, v, len(p))
	}
	return newbook
}

func load() map[string]string {
	if len(os.Args) < 2 {
		panic("Must pass input file on commandline.")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rules := map[string]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), " => ")

		rules[ss[0]] = ss[1]
	}
	return rules
}

func makeGrid(size int) pixels {
	g := make(pixels, 0, size)
	for i := 0; i < size; i++ {
		g = append(g, make([]byte, size))
	}
	return g
}

func enhance(p pixels, rulebook map[string]string) pixels {
	var newp pixels
	var g, groups int
	if len(p)%2 == 0 {
		// squares of 2
		newp = makeGrid(3 * (len(p) / 2))
		g = 2
		groups = len(p) / 2
	} else {
		// squares of 3
		newp = makeGrid(4 * (len(p) / 3))
		g = 3
		groups = len(p) / 3
	}

	sig := bytes.NewBuffer(make([]byte, g*g+(g-1)))
	for i := 0; i < groups; i++ {
		for j := 0; j < groups; j++ {
			// Create the lookup signature
			sig.Reset()
			for a := 0; a < g; a++ {
				for b := 0; b < g; b++ {
					sig.WriteByte(p[i*g+a][j*g+b])
				}
				if a != g-1 {
					sig.WriteByte('/')
				}
			}
			// Find the enhanced pattern
			r, ok := rulebook[sig.String()]
			if !ok {
				panic("Unknown pattern " + sig.String())
			}
			// Copy the enhanced pattern into the grid
			for a := 0; a < g+1; a++ {
				for b := 0; b < g+1; b++ {
					newp[i*(g+1)+a][j*(g+1)+b] = r[a*(g+2)+b]
				}
			}
		}
	}
	return newp
}

func Solve() {
	rulebook := load()
	rulebook = expand(rulebook)

	p := makeGrid(3)
	p.fill(".#./..#/###")

	for i := 0; i < 18; i++ {
		p = enhance(p, rulebook)
		if i == 4 {
			fmt.Println("Part A:", p.OnCount())
		}
		if i == 17 {
			fmt.Println("Part B:", p.OnCount())
		}
	}
}
