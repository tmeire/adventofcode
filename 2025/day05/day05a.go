package day05

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func merge(rl []r) []r {
	var rn []r

	n := rl[0]
	for i := 1; i < len(rl); i++ {
		if rl[i].l <= n.h+1 {
			n.h = max(n.h, rl[i].h)
		} else {
			rn = append(rn, n)
			n = rl[i]
		}
	}
	return append(rn, n)
}

func Solve() {
	rl, il := read("./2025/day05/input.txt")

	sort.Slice(rl, func(i, j int) bool { return rl[i].l < rl[j].l })

	rl = merge(rl)

	n := 0
	for _, i := range il {
		for _, r := range rl {
			if i >= r.l && i <= r.h {
				n++
				break
			}
		}
	}
	println(n)

	nt := 0
	for _, r := range rl {
		nt += r.h - r.l + 1
	}
	println(nt)
}

type r struct {
	l, h int
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func read(fn string) ([]r, []int) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var rl []r
	var i []int
	inRanges := true

	bf := bufio.NewReader(f)
	l, _, err := bf.ReadLine()
	for err == nil {
		//println(string(l))
		if inRanges {
			if len(l) == 0 {
				inRanges = false
				l, _, err = bf.ReadLine()
				continue
			}
			var low, high int
			_, err := fmt.Sscanf(string(l), "%d-%d", &low, &high)
			if err != nil {
				panic(err)
			}
			rl = append(rl, r{low, high})
		} else {
			i = append(i, atoi(string(l)))
		}
		l, _, err = bf.ReadLine()
	}
	if !errors.Is(err, io.EOF) {
		panic(err.Error())
	}
	return rl, i
}
