package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tmeire/adventofcode/io"
)

type chemical struct {
	name   string
	amount int
}

type reaction struct {
	amount int // amount generated
	inputs map[string]int
}

func parseChemical(s string) (string, int) {
	sp := strings.Split(s, " ")
	i, err := strconv.Atoi(sp[0])
	if err != nil {
		panic(err)
	}
	return sp[1], i
}

func unfold(chemicals map[string]int, reactions map[string]reaction, r reaction) {
	for name, amount := range r.inputs {
		chemicals[name] += amount
		if name != "ORE" {
			reaction := reactions[name]
			//fmt.Printf("%#v\n", chemicals)
			for chemicals[name] >= reaction.amount {
				//fmt.Printf("%#v\n", chemicals)
				chemicals[name] -= reaction.amount
				unfold(chemicals, reactions, reaction)
				//fmt.Printf("%#v\n", chemicals)
			}
		}
	}
}

func calculateOreAmount(reactions map[string]reaction, units int) int {
	want := []chemical{chemical{"FUEL", units}}
	have := make(map[string]int)
	done := make(map[string]int)

	for len(want) > 0 {
		wantchem := want[0]
		if have[wantchem.name] > 0 {
			used := min(wantchem.amount, have[wantchem.name])
			have[wantchem.name] -= used
			wantchem.amount -= used
		}

		if wantchem.amount > 0 {
			reaction := reactions[wantchem.name]
			_, needsORE := reaction.inputs["ORE"]
			if !needsORE {
				iter := wantchem.amount / reaction.amount
				if wantchem.amount%reaction.amount != 0 {
					iter++
				}

				for name, amount := range reaction.inputs {
					want = append(want, chemical{name, iter * amount})
				}
				have[wantchem.name] += (iter * reaction.amount) - wantchem.amount
			} else {
				done[wantchem.name] += wantchem.amount
			}
		}

		want = want[1:len(want)]
	}

	ore := 0
	for name, amount := range done {
		reaction := reactions[name]

		iter := amount / reaction.amount
		if amount%reaction.amount != 0 {
			iter++
		}
		ore += iter * reaction.inputs["ORE"]
	}
	return ore
}

func main() {
	reactions := make(map[string]reaction)

	data, err := io.ReadLinesFromFile("data.txt")
	if err != nil {
		panic(err.Error())
	}
	for _, l := range data {
		parts := strings.Split(l, " => ")
		outn, outi := parseChemical(parts[1])

		r := reaction{outi, make(map[string]int)}
		for _, chem := range strings.Split(parts[0], ", ") {
			outn, outi := parseChemical(chem)
			r.inputs[outn] = outi
		}
		reactions[outn] = r
	}

	ore := calculateOreAmount(reactions, 1)
	fmt.Printf("1 FUEL = %d ORE", ore)

	x := 1000000000000
	units := 2
	for ore < x {
		units = units * units
		ore = calculateOreAmount(reactions, units)
		println(units, ore)
	}
	if ore == x {
		println(units)
		return
	}
	min := 1
	max := units
	for max > min+1 {
		mid := (max + min) / 2
		ore = calculateOreAmount(reactions, mid)
		println(min, mid, max, ore)
		if ore == x {
			println(mid)
			return
		}
		if ore < x {
			min = mid
		}
		if ore > x {
			max = mid
		}
	}

	println(min)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
