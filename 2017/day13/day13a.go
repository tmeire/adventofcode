package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Layer interface {
	Jump(delay int)
	Move()
	Severity() int
	Caught() bool
	Reset()
}

type ZeroLayer struct{}

func (l *ZeroLayer) Jump(delay int) {
	// noop
}

func (l *ZeroLayer) Move() {
	// noop
}

func (l ZeroLayer) Severity() int {
	return 0
}

func (l ZeroLayer) Caught() bool {
	return false
}

func (l *ZeroLayer) Reset() {
	// noop
}

type SecurityLayer struct {
	Depth int
	Range int

	Scanner int
	Dir     bool
}

func (l *SecurityLayer) Jump(delay int) {
	moves := delay % (2 * (l.Range - 1))
	if moves < l.Range {
		l.Scanner = moves
		l.Dir = (l.Scanner < l.Range-1)
	} else {
		l.Scanner = 2*(l.Range-1) - moves
		l.Dir = (l.Scanner == 0)
	}
}

func (l *SecurityLayer) Move() {
	if l.Dir {
		l.Scanner++
		// if scanner is at the end, we move back down again
		l.Dir = (l.Scanner < l.Range-1)
	} else {
		l.Scanner--
		// if scanner is 0, we move back up again
		l.Dir = (l.Scanner == 0)
	}
}

func (l SecurityLayer) Severity() int {
	return l.Depth * l.Range
}

func (l SecurityLayer) Caught() bool {
	return l.Scanner == 0
}

func (l *SecurityLayer) Reset() {
	l.Scanner = 0
	l.Dir = true
}

type Firewall []Layer

func (f Firewall) Passthrough() int {
	sev := 0

	// Pass through the firewall
	pos := -1
	for i := 0; i < len(f); i++ {
		// Move the packet forward
		pos++
		// Check if caught
		if f[pos].Caught() {
			//fmt.Printf("Caught at layer %d, sev = %d, total: %d\n", pos, f[pos].Severity(), sev)
			sev += f[pos].Severity()
		}
		// Move the scanners forward
		for _, l := range f {
			l.Move()
		}
	}
	return sev
}

func (f Firewall) Caught(delay int) bool {
	// Move the scanners forward for `delay` picoseconds.
	for _, l := range f {
		l.Jump(delay)
	}

	// Pass through the firewall
	pos := -1
	for i := 0; i < len(f); i++ {
		// Move the packet forward
		pos++
		// Check if caught
		if f[pos].Caught() {
			return true
		}
		// Move the scanners forward
		for _, l := range f {
			l.Move()
		}
	}
	return false
}

func (f Firewall) Reset() {
	for _, l := range f {
		l.Reset()
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func load(fname string) []Layer {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	layers := []Layer{}
	prevdepth := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		o := strings.Split(scanner.Text(), ": ")
		sl := &SecurityLayer{atoi(o[0]), atoi(o[1]), 0, true}

		// Add in a zero layer on indeces that are not in the input
		for prevdepth < sl.Depth {
			layers = append(layers, &ZeroLayer{})
			prevdepth++
		}

		layers = append(layers, sl)
		prevdepth++
	}
	return layers
}

func Solve() {
	if len(os.Args) < 2 {
		fmt.Println("Expected the input file name as commandline argument.")
		os.Exit(1)
	}

	fw := Firewall(load(os.Args[1]))

	fmt.Println("Part A:", fw.Passthrough())

	fw.Reset()
	delay := 1
	for fw.Caught(delay) {
		fw.Reset()
		delay++
	}
	fmt.Printf("Delay %d not caught\n", delay)
	fmt.Println("Part B:", delay)
}
