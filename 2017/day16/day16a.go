package day16

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ProgramList map[string]int

func (pl ProgramList) Ordered() []string {
	ordered := make([]string, len(pl)+1)
	for k, v := range pl {
		ordered[v] = k
	}
	return ordered
}

func (pl ProgramList) String() string {
	return strings.Join(pl.Ordered(), "")
}

type Mover interface {
	move(ProgramList)
}

type Shifter struct {
	a int
}

func (s Shifter) move(pl ProgramList) {
	for k, v := range pl {
		pl[k] = (v + s.a) % len(pl)
	}
}

type Exchanger struct {
	a, b int
}

func (e Exchanger) move(pl ProgramList) {
	var x, y string
	for k, v := range pl {
		if v == e.a {
			x = k
		}
		if v == e.b {
			y = k
		}
	}
	pl[x] = e.b
	pl[y] = e.a
}

type Partner struct {
	a, b string
}

func (p Partner) move(pl ProgramList) {
	pl[p.a], pl[p.b] = pl[p.b], pl[p.a]
}

func load(fname string) []Mover {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	moves := strings.Split(strings.Replace(string(content), "/", " ", -1), ",")

	movers := make([]Mover, 0, len(moves))
	for _, m := range moves {
		switch m[0] {
		case 's':
			sh := Shifter{}
			fmt.Sscan(m[1:], &(sh.a))
			movers = append(movers, sh)
		case 'x':
			ex := Exchanger{}
			fmt.Sscanf(m[1:], "%d %d", &(ex.a), &(ex.b))
			movers = append(movers, ex)
		case 'p':
			p := Partner{}
			fmt.Sscanf(m[1:], "%s %s", &(p.a), &(p.b))
			movers = append(movers, p)
		}
	}
	return movers
}

func Solve() {
	if len(os.Args) < 2 {
		fmt.Println("Expected the input file name as commandline argument.")
		os.Exit(1)
	}

	movers := load(os.Args[1])

	pl := make(ProgramList)
	for i, v := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"} {
		pl[v] = i
	}

	for _, m := range movers {
		m.move(pl)
	}
	fmt.Println("Part A:", pl.String())

	positions := make(ProgramList)
	positions[pl.String()] = 1

	var iterlen int
	for i := 1; i < 1000000000 && iterlen == 0; i++ {
		for _, m := range movers {
			m.move(pl)
		}

		s := pl.String()
		if j, ok := positions[s]; ok {
			iterlen = i + 1 - j
		} else {
			positions[s] = i + 1
		}
	}
	ordered := positions.Ordered()

	fmt.Println("Part B:", ordered[1000000000%iterlen])
}
