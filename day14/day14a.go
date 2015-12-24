package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Reindeer struct {
	speed  int // add which speed does it travel
	travel int // how long can it travel
	rest   int // how long must it rest after travel
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (r *Reindeer) Travel(s int) int {
	t := s / (r.travel + r.rest)

	// distance in the normal timeframe
	dist := (t * r.travel * r.speed)

	// remaining seconds or r.travel if too many remaining
	rem := min(r.travel, s-(r.travel+r.speed)*t)

	return dist + rem*r.speed
}

var regex = regexp.MustCompile("^.* can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.")

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func parse(t string) *Reindeer {
	m := regex.FindStringSubmatch(t)
	return &Reindeer{atoi(m[1]), atoi(m[2]), atoi(m[3])}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := make([]*Reindeer, 0, 9)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		r = append(r, parse(t))
	}

	best := 0
	for _, d := range r {
		dist := d.Travel(2503)
		if best < dist {
			best = dist
		}
	}
	fmt.Println(best)
}
