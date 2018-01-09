package day22

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

func (d Direction) left() Direction {
	if d == 0 {
		return 3
	}
	return d - 1
}

func (d Direction) right() Direction {
	if d == 3 {
		return 0
	}
	return d + 1
}

const (
	UP    Direction = 0
	RIGHT           = 1
	DOWN            = 2
	LEFT            = 3
)

func move(x, y int, d Direction) (int, int) {
	//fmt.Println(x, y, d)
	switch d {
	case UP:
		return x - 1, y
	case RIGHT:
		return x, y + 1
	case DOWN:
		return x + 1, y
	case LEFT:
		return x, y - 1
	}
	panic("Unknown direction")
}

type Infection int

func (i Infection) evolve(rate int) Infection {
	return Infection((int(i) + rate) % 4)
}

const (
	CLEAN    Infection = 0
	WEAKENED           = 1
	INFECTED           = 2
	FLAGGED            = 3
)

type grid map[int]map[int]Infection

func (g grid) set(x, y int, v Infection) {
	r, ok := g[x]
	if !ok {
		r = make(map[int]Infection)
		g[x] = r
	}
	r[y] = v
}

func (g grid) print() {
	fmt.Println("====")
	for i := -4; i <= 4; i++ {
		for j := -4; j <= 4; j++ {
			if g[i][j] == INFECTED {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("====")
}

func load() grid {
	if len(os.Args) < 2 {
		panic("Must pass input file on commandline.")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	g := make(grid)

	input := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	x := len(input[0]) / 2
	y := len(input) / 2

	for i, s := range input {
		for j, n := range s {
			infected := CLEAN
			if n == '#' {
				infected = INFECTED
			}
			g.set(i-y, j-x, infected)
		}
	}
	return g
}

func infect(bursts int, evolved bool) int {
	rate := 2
	if evolved {
		rate = 1
	}

	g := load()

	dir := UP

	x, y := 0, 0

	infected := 0
	for i := 0; i < bursts; i++ {
		if g[x][y] == INFECTED {
			dir = dir.right()
		} else if g[x][y] == FLAGGED {
			// Turn round (double right)
			dir = dir.right().right()
		} else if g[x][y] == WEAKENED {
			// Does not turn, but will result in infection for part B
			if evolved {
				infected++
			}
		} else {
			dir = dir.left()
			// Will result in infection for part A
			if !evolved {
				infected++
			}
		}
		g.set(x, y, g[x][y].evolve(rate))
		x, y = move(x, y, dir)
	}
	return infected
}

func Solve() {
	fmt.Println("Part A:", infect(10000, false))
	fmt.Println("Part B:", infect(10000000, true))

}
