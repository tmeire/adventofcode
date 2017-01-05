package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parse(s string) *ingredient {
	var i ingredient
	fmt.Sscanf(s, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &(i.name), &(i.capacity), &(i.durability), &(i.flavor), &(i.texture), &(i.calories))
	return &i
}

func score(ingr []*ingredient, dist []int) int {
	capa := 0
	for i, x := range ingr {
		capa += dist[i] * x.capacity
	}
	if capa < 0 {
		capa = 0
	}

	dur := 0
	for i, x := range ingr {
		dur += dist[i] * x.durability
	}
	if dur < 0 {
		dur = 0
	}

	flav := 0
	for i, x := range ingr {
		flav += dist[i] * x.flavor
	}
	if flav < 0 {
		flav = 0
	}

	tex := 0
	for i, x := range ingr {
		tex += dist[i] * x.texture
	}
	if tex < 0 {
		tex = 0
	}

	return capa * dur * flav * tex
}

func dists() [][]int {
	d := make([][]int, 0)
	for i := 0; i <= 100; i += 1 {
		for j := 0; i+j <= 100; j += 1 {
			for k := 0; i+j+k <= 100; k += 1 {
				d = append(d, []int{i, j, k, 100 - i - j - k})
			}
		}
	}
	return d
}

func optmize(i []*ingredient) int {
	ds := dists()

	best := 0
	for _, dist := range ds {
		x := score(i, dist)
		if x > best {
			best = x
		}
	}
	return best
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ingr := make([]*ingredient, 0, 4)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ingr = append(ingr, parse(scanner.Text()))
	}
	fmt.Println(optmize(ingr))
}
