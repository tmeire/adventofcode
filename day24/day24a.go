package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func sum(v []int) int {
	s := 0
	for _, i := range v {
		s += i
	}
	return s
}

func mini(a int64, b int64) int64 {
	if a != -1 && b == -1 {
		return a
	}
	if a == -1 && b != -1 {
		return b
	}
	if a < b {
		return a
	} else {
		return b
	}
}

func group(packages []int64, idx int, target int64, s1 int64, s2 int64, s3 int64, qe int64, bestQE int64) int64 {
	if s1 > target || s2 > target || s3 > target || qe > bestQE {
		return -1
	}

	if idx == len(packages) {
		if s1 == target && s2 == target && s3 == target {
			return qe
		} else {
			return -1
		}
	} else {
		qe1 := group(packages, idx+1, target, s1+packages[idx], s2, s3, qe*packages[idx], bestQE)
		if qe1 != -1 && qe1 < bestQE {
			bestQE = qe1
		}
		qe2 := group(packages, idx+1, target, s1, s2+packages[idx], s3, qe, bestQE)
		if qe2 != -1 && qe2 < bestQE {
			bestQE = qe2
		}
		qe3 := group(packages, idx+1, target, s1, s2, s3+packages[idx], qe, bestQE)
		return mini(mini(qe1, qe2), qe3)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	packages := make([]int, 0, 28)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		packages = append(packages, n)
	}

	s := sum(packages)
	sort.Sort(sort.Reverse(sort.IntSlice(packages)))

	spg := s / 3
	fmt.Println(s, spg)

	p64 := make([]int64, len(packages))
	for i, x := range packages {
		p64[i] = int64(x)
	}

	fmt.Println(group(p64, 0, int64(spg), 0, 0, 0, 1, math.MaxInt64))
}
