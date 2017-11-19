package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type disc struct {
	positions int
	start     int
}

func (d disc) positioned(time int) bool {
	return ((d.start + time) % d.positions) == 0
}

func check(discs []disc, time int) bool {
	for i, d := range discs {
		if !d.positioned(time + i + 1) {
			return false
		}
	}
	return true
}

var REGEX = regexp.MustCompile(`Disc #[0-9]+ has ([0-9]+) positions; at time=0, it is at position ([0-9]+).`)

func parse(s string) disc {
	ss := REGEX.FindStringSubmatch(s)

	pos, _ := strconv.Atoi(ss[1])
	start, _ := strconv.Atoi(ss[2])

	return disc{pos, start}
}

func load() []disc {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	discs := []disc{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		discs = append(discs, parse(scanner.Text()))
	}
	return discs
}

func main() {
	//Disc #1 has 5 positions; at time=0, it is at position 4.
	//Disc #2 has 2 positions; at time=0, it is at position 1.
	//discs := []disc{
	//	{5, 4},
	//	{2, 1},
	//}

	discs := load()

	i := 0
	for !check(discs, i) {
		i += 1
	}
	fmt.Println("Capsule dropped at", i)

	// Part b, with an extra disc
	discs = append(discs, disc{11, 0})

	i = 0
	for !check(discs, i) {
		i += 1
	}
	fmt.Println("Capsule dropped at", i)
}
