package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile(".* to .* = ([0-9]*)")

func parseDistance(t string) int {
	m := regex.FindStringSubmatch(t)
	i, err := strconv.Atoi(m[1])
	if err != nil {
		panic("Invalid distance " + m[1])
	}
	return i
}

func computeAllRoutes() [][]int {
	routes := make([][]int, 40320)
	i := 0
	for a := 0; a < 8; a += 1 {
		for b := 0; b < 8; b += 1 {
			if b == a {
				continue
			}
			for c := 0; c < 8; c += 1 {
				if c == a || c == b {
					continue
				}
				for d := 0; d < 8; d += 1 {
					if d == a || d == b || d == c {
						continue
					}
					for e := 0; e < 8; e += 1 {
						if e == a || e == b || e == c || e == d {
							continue
						}
						for f := 0; f < 8; f += 1 {
							if f == a || f == b || f == c || f == d || f == e {
								continue
							}
							for g := 0; g < 8; g += 1 {
								if g == a || g == b || g == c || g == d || g == e || g == f {
									continue
								}
								for h := 0; h < 8; h += 1 {
									if h == a || h == b || h == c || h == d || h == e || h == f || h == g {
										continue
									}
									route := make([]int, 8)
									route[0] = a
									route[1] = b
									route[2] = c
									route[3] = d
									route[4] = e
									route[5] = f
									route[6] = g
									route[7] = h
									routes[i] = route
									i += 1
								}
							}
						}
					}
				}
			}
		}
	}
	return routes
}

func computeCost(route []int) int {
	cost := 0
	for i := 0; i < len(route)-1; i += 1 {
		cost += distances[route[i]][route[i+1]]
	}
	return cost
}

var distances [8][8]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	i := 0
	j := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		distances[i][j] = parseDistance(t)
		distances[j][i] = distances[i][j]

		j += 1
		if j == 8 {
			i += 1
			j = i + 1
		}
	}

	routes := computeAllRoutes()

	max := 0
	for _, r := range routes {
		c := computeCost(r)
		if c > max {
			max = c
		}
	}
	fmt.Printf("%d\n", max)
}
