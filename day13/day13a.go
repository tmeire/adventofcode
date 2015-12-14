package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

var r = regexp.MustCompile("^([a-zA-Z]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([a-zA-Z]+).$")

func parse(s string) (string, string, int) {
	m := r.FindStringSubmatch(s)

	score, _ := strconv.Atoi(m[3])
	if m[2] == "lose" {
		score *= -1
	}

	return m[1], m[4], score
}

func computeAllArrangements() [][]int {
	arrs := make([][]int, 40320)
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
									arr := make([]int, 8)
									arr[0] = a
									arr[1] = b
									arr[2] = c
									arr[3] = d
									arr[4] = e
									arr[5] = f
									arr[6] = g
									arr[7] = h
									arrs[i] = arr
									i += 1
								}
							}
						}
					}
				}
			}
		}
	}
	return arrs
}

func computeScoreForArrangenment(arr []int) int {
	score := scores[names[arr[0]]][names[arr[1]]] + scores[names[arr[0]]][names[arr[len(arr)-1]]]

	for i := 1; i < len(arr)-1; i += 1 {
		score += scores[names[arr[i]]][names[arr[i+1]]] + scores[names[arr[i]]][names[arr[i-1]]]
	}
	score += scores[names[arr[len(arr)-1]]][names[arr[0]]] + scores[names[arr[len(arr)-1]]][names[arr[len(arr)-2]]]
	return score
}

var scores map[string]map[string]int
var names []string

func main() {
	scores = make(map[string]map[string]int)
	names = make([]string, 0, 8)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	name := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a, b, s := parse(scanner.Text())

		// store the name
		if name != a {
			names = append(names, a)
			name = a
		}

		// store the score
		if _, ok := scores[a]; !ok {
			scores[a] = make(map[string]int)
		}
		scores[a][b] = s
	}
	fmt.Println(scores)
	fmt.Println(names)

	best := math.MinInt32
	for _, arr := range computeAllArrangements() {
		score := computeScoreForArrangenment(arr)
		if score > best {
			best = score
		}
	}
	fmt.Println(best)
}
