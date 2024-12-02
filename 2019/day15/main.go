package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/tmeire/adventofcode/intcode"
)

type grid struct {
	d map[int64]map[int64]int64

	minx, maxx, miny, maxy int64
}

func (g *grid) get(i, j int64) int64 {
	x, ok := g.d[i][j]
	if !ok {
		return -1
	}
	return x
}

func (g *grid) set(i, j, v int64) {
	r, ok := g.d[i]
	if !ok {
		r = make(map[int64]int64)
		g.d[i] = r
	}
	g.d[i][j] = v

	g.screensize(i, j)
}

func (g *grid) screensize(i, j int64) {
	if g.minx > i {
		g.minx = i
	}
	if g.maxx < i {
		g.maxx = i
	}
	if g.miny > j {
		g.miny = j
	}
	if g.maxy < j {
		g.maxy = j
	}
}

func (g *grid) getChar(i, j int64) rune {
	p := g.get(i, j)
	switch p {
	case 0:
		return '▓'
	case 1:
		return '.'
	case 2:
		return 'X'
	}
	return ' '
}

func (g *grid) paint() {
	println(g.minx, g.maxx, g.miny, g.maxy)
	for j := g.miny; j <= g.maxy; j++ {
		for i := g.minx; i <= g.maxx; i++ {
			print(string(g.getChar(i, j)))
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

type direction int64

const (
	NORTH direction = 1
	SOUTH           = 2
	WEST            = 3
	EAST            = 4
)

func (d direction) turn(r direction) direction {
	// 1 left, 2 straight, 3 right, back
	switch r {
	case 1:
		// TURN 90 DEG COUNTER CLOCK
		switch d {
		case NORTH:
			return WEST
		case SOUTH:
			return EAST
		case WEST:
			return SOUTH
		case EAST:
			return NORTH
		}
	case 2:
		// TURN 180 DEG COUNTER CLOCK
		switch d {
		case NORTH:
			return SOUTH
		case SOUTH:
			return NORTH
		case WEST:
			return EAST
		case EAST:
			return WEST
		}
	case 3:
		// TURN 270 DEG COUNTER CLOCK
		switch d {
		case NORTH:
			return EAST
		case SOUTH:
			return WEST
		case WEST:
			return NORTH
		case EAST:
			return SOUTH
		}
	}
	// FOLLOW ALONG THE SAME DIRECTION
	return d
}

const (
	UNKNOWN int64 = -1
	WALL          = 0
	OPEN          = 1
	OXYGEN        = 2
)

func main() {
	data := fromFile()

	p := intcode.NewProgram(data)

	go p.Simulate()

	floor := &grid{d: make(map[int64]map[int64]int64)}
	steps := &grid{d: make(map[int64]map[int64]int64)}

	//moves := make([]int, 0)

	done := false
	var v direction = NORTH
	var x, y, oxyX, oxyY int64
	floor.set(0, 0, OPEN)
	i := 0
	for !done {
		select {
		case p.Stdin <- int64(v):
		case out := <-p.Stdout:
			// Droid was able to move to the position (not a wall), so update x,y to the right position
			switch v {
			case NORTH:
				floor.set(x, y-1, out)
				if out != 0 {
					y--
				}
			case EAST:
				floor.set(x+1, y, out)
				if out != 0 {
					x++
				}
			case SOUTH:
				floor.set(x, y+1, out)
				if out != 0 {
					y++
				}
			case WEST:
				floor.set(x-1, y, out)
				if out != 0 {
					x--
				}
			}

			if out == OXYGEN {
				oxyX, oxyY = x, y
				println(oxyX, oxyY)
				done = true
			}
			// so where does the droid go next?
			if floor.get(coord(x, y, v.turn(1))) != WALL {
				v = v.turn(1)
			}
			if floor.get(coord(x, y, v.turn(3))) != WALL {
				v = v.turn(3)
			}
			if floor.get(coord(x, y, v.turn(4))) != WALL {
				v = v.turn(4)
			}
			if floor.get(coord(x, y, v.turn(2))) != WALL {
				v = v.turn(2)
			}
			//println(v)

			println("----")
			floor.paint()
			println("----")
			if i == 100 {
				done = true
			}
			i++
		case <-p.Done:
			println("DONE")
			done = true
		}
	}
	println("min steps:", steps.get(oxyX, oxyY))

}

func coord(x, y int64, v direction) (int64, int64) {
	switch v {
	case NORTH:
		return x, y - 1
	case EAST:
		return x + 1, y
	case SOUTH:
		return x, y + 1
	case WEST:
		return x - 1, y
	}
	return 0, 0
}

func min(ints ...int64) int64 {
	if len(ints) == 1 {
		return ints[0]
	}
	a := ints[0]
	b := min(ints[1:]...)

	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a < b {
		return a
	}
	return b
}
