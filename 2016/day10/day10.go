package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type recipient interface {
	put(chip int)
}

type robot struct {
	id string

	rlow, rhigh recipient

	chip1, chip2 int
}

func (r *robot) put(chip int) {
	if r.chip1 == -1 {
		r.chip1 = chip
	} else {
		r.chip2 = chip

		if (r.chip1 == 61 && r.chip2 == 17) || (r.chip1 == 17 && r.chip2 == 61) {
			fmt.Println("Robot " + r.id + " is responsible to compare 61 and 17")
		}
		r.forward()
	}
}

func (r *robot) forward() {
	r.rlow.put(r.low())
	r.rhigh.put(r.high())
}

func (r *robot) low() int {
	if r.chip1 < r.chip2 {
		return r.chip1
	}
	return r.chip2
}

func (r *robot) high() int {
	if r.chip1 > r.chip2 {
		return r.chip1
	}
	return r.chip2
}

type bin struct {
	id   string
	chip int
}

func (b *bin) put(chip int) {
	b.chip = chip
}

type factory struct {
	bots map[string]*robot
	bins map[string]*bin
}

func (f *factory) get(t string, n string) recipient {
	switch t {
	case "bot":
		b, ok := f.bots[n]
		if ok {
			return b
		}
		b = &robot{n, nil, nil, -1, -1}
		f.bots[n] = b
		return b
	case "output":
		b, ok := f.bins[n]
		if ok {
			return b
		}
		b = &bin{n, -1}
		f.bins[n] = b
		return b
	}
	panic("Unknown recipient type " + t)
}

func (f *factory) binin(v int, n string) {
	b := f.get("bot", n)

	b.put(v)
}

func (f *factory) pass(n, tlow, nlow, thigh, nhigh string) {
	b := f.get("bot", n).(*robot)
	b.rlow = f.get(tlow, nlow)
	b.rhigh = f.get(thigh, nhigh)
}

var INPUT_REGEX = regexp.MustCompile(`value ([0-9]+) goes to bot ([0-9]+)`)
var PASS_REGEX = regexp.MustCompile(`bot ([0-9]+) gives low to (bot|output) ([0-9]+) and high to (bot|output) ([0-9]+)`)

func main() {
	f := &factory{make(map[string]*robot), make(map[string]*bin)}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cmds := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmds = append(cmds, scanner.Text())
	}

	// sort them to make sure the command that build the network come first
	sort.Strings(cmds)

	for _, cmd := range cmds {
		p := INPUT_REGEX.FindStringSubmatch(cmd)
		if len(p) == 3 {
			i, err := strconv.ParseInt(p[1], 10, 64)
			if err != nil {
				panic(err)
			}
			f.binin(int(i), p[2])
			continue
		}

		p = PASS_REGEX.FindStringSubmatch(cmd)
		if len(p) == 6 {
			f.pass(p[1], p[2], p[3], p[4], p[5])
			continue
		}

		panic("BAD COMMAND: " + cmd)
	}

	// part b of day 10
	b0 := f.get("output", "0").(*bin)
	b1 := f.get("output", "1").(*bin)
	b2 := f.get("output", "2").(*bin)

	fmt.Println("The value of the three bins is ", (b0.chip)*(b1.chip)*(b2.chip))
}
