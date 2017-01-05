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

func max(a []int) int {
	m := 0
	for _, b := range a {
		if b > m {
			m = b
		}
	}
	return m
}

func (r *Reindeer) Travel(s int) <-chan int {
	c := make(chan int)
	go func() {
		dist := 0
		i := 0
		for i < s {
			for a := 0; a < r.travel && i < s; a += 1 {
				dist += r.speed
				c <- dist
			}
			for a := 0; a < r.rest && i < s; a += 1 {
				c <- dist
			}
		}
		close(c)
	}()
	return c
}

var regex = regexp.MustCompile("^.* can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.")

const MAX_SECONDS = 2503

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func parse(t string) <-chan int {
	m := regex.FindStringSubmatch(t)
	r := &Reindeer{atoi(m[1]), atoi(m[2]), atoi(m[3])}
	return r.Travel(MAX_SECONDS)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reindeers := make([]<-chan int, 0, 9)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		reindeers = append(reindeers, parse(t))
	}

	scores := make([]int, len(reindeers))

	var progress [9]int
	for s := 0; s < MAX_SECONDS; s += 1 {
		mx := 0
		for i, r := range reindeers {
			progress[i] = <-r
			if mx < progress[i] {
				mx = progress[i]
			}
		}

		for i, p := range progress {
			if p == mx {
				scores[i] += 1
			}
		}
	}
	fmt.Println(max(scores))
}
