package day24

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type component struct {
	a, b int
	used bool
}

func Strongest(comps []*component, pin, strength, length int, longest bool) (int, int) {
	maxStrength := strength
	maxLength := length
	for _, c := range comps {
		if c.used {
			continue
		}
		if c.a == pin {
			c.used = true
			sn, ln := Strongest(comps, c.b, strength+c.a+c.b, length+1, longest)
			if sn > maxStrength && (!longest || ln > maxLength) {
				maxStrength = sn
				maxLength = ln
			}
			c.used = false
		} else if c.b == pin {
			c.used = true
			sn, ln := Strongest(comps, c.a, strength+c.a+c.b, length+1, longest)
			if sn > maxStrength && (!longest || ln > maxLength) {
				maxStrength = sn
				maxLength = ln
			}
			c.used = false
		}
	}
	return maxStrength, maxLength
}

func atoi(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}

func load() []*component {
	if len(os.Args) < 2 {
		panic("Must pass input file on commandline.")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	comps := []*component{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var c component

		ss := strings.Split(scanner.Text(), "/")

		c.a = atoi(ss[0])
		c.b = atoi(ss[1])
		comps = append(comps, &c)
	}
	return comps
}

func Solve() {
	comps := load()
	maxStrength, _ := Strongest(comps, 0, 0, 0, false)
	fmt.Println("Part A", maxStrength)
	maxStrength, maxLength := Strongest(comps, 0, 0, 0, true)
	fmt.Println("Part B", maxStrength, maxLength)
}
