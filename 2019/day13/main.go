package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/tmeire/adventofcode/intcode"
)

type grid struct {
	d map[int64]map[int64]rune

	minx, maxx, miny, maxy int64
}

func (g *grid) get(i, j int64) rune {
	p := g.d[i][j]
	if p == 0 {
		return ' '
	}
	return p
}

func (g *grid) set(i, j int64, v int64) {
	r, ok := g.d[i]
	if !ok {
		r = make(map[int64]rune)
		g.d[i] = r
	}
	switch v {
	case 0:
		g.d[i][j] = ' '
	case 1:
		g.d[i][j] = '▓'
	case 2:
		g.d[i][j] = '#'
	case 3:
		g.d[i][j] = '_'
	case 4:
		g.d[i][j] = '●'
	}
	g.screensize(i, j)
}

func (g *grid) screensize(i, j int64) {
	if g.minx > i {
		g.minx = i
	}
	if g.maxx < i {
		g.maxx = i
	}
	if g.miny > i {
		g.miny = i
	}
	if g.maxy < i {
		g.maxy = i
	}
}

func (g *grid) paint() {
	for j := g.miny; j <= g.maxy; j++ {
		for i := g.minx; i <= g.maxx; i++ {
			print(string(g.get(i, j)))
		}
		println()
	}
}

func fromFile() []int64 {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err.Error())
	}

	intchar := strings.Split(string(b), ",")
	data := make([]int64, len(intchar))
	for x, chars := range intchar {
		i, err := strconv.ParseInt(chars, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		data[x] = i
	}
	return data
}

func main() {
	data := fromFile()

	p := intcode.NewProgram(data)

	go p.Simulate()

	var walls int64
	done := false
	for !done {
		select {
		case x := <-p.Stdout:
			x = <-p.Stdout
			x = <-p.Stdout
			if x == 2 {
				walls++
			}
		case <-p.Done:
			done = true
		}
	}
	println("Walls", walls)

	data = fromFile()
	// Put in coins
	data[0] = 2

	screen := &grid{d: make(map[int64]map[int64]rune)}

	p = intcode.NewProgram(data)

	go p.Simulate()

	done = false

	var score, balx, padx, s int64

	for !done {
		select {
		case x := <-p.Stdout:
			y := <-p.Stdout
			v := <-p.Stdout
			if x != -1 {
				screen.set(x, y, v)
				if v == 4 {
					balx = x
					s = sign(balx - padx)
				} else if v == 3 {
					padx = x
					s = sign(balx - padx)
				}
				//screen.paint()
			} else {
				score = v
			}
		case p.Stdin <- s:
		case <-p.Done:
			done = true
		}
	}
	println("Final score:", score)
}

func sign(v int64) int64 {
	if v < 0 {
		return -1
	}
	if v > 0 {
		return 1
	}
	return 0
}
