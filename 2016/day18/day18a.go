package main

import "fmt"

const INPUT = `...^^^^^..^...^...^^^^^^...^.^^^.^.^.^^.^^^.....^.^^^...^^^^^^.....^.^^...^^^^^...^.^^^.^^......^^^^`

func isSafe(a, b, c bool) bool {
	return !((!a && !b && c) || (a && !b && !c) || (!a && b && c) || (a && b && !c))
}

func genNext(tiles []bool) []bool {
	next := make([]bool, len(tiles))

	next[0] = isSafe(true, tiles[0], tiles[1])

	for i := 1; i < len(tiles)-1; i++ {
		next[i] = isSafe(tiles[i-1], tiles[i], tiles[i+1])
	}
	next[len(tiles)-1] = isSafe(tiles[len(tiles)-2], tiles[len(tiles)-1], true)
	return next
}

func load(s string) []bool {
	tiles := []bool{}
	fmt.Println(len(s))
	for _, x := range s {
		switch x {
		case '.':
			tiles = append(tiles, true)
		case '^':
			tiles = append(tiles, false)
		}
	}
	return tiles
}

func print(b []bool) {
	for _, s := range b {
		if s {
			fmt.Print(".")
		} else {
			fmt.Print("^")
		}
	}
	fmt.Println()
}

func count(s string, rows int) int {
	tiles := load(s)

	safes := 0
	for i := 0; i < rows; i++ {
		//fmt.Print(i, " ")
		//print(tiles)
		for _, x := range tiles {
			if x {
				safes++
			}
		}
		tiles = genNext(tiles)
	}
	return safes
}

func main() {
	fmt.Println(count("..^^.", 3))
	fmt.Println(count(".^^.^.^^^^", 10))
	fmt.Println(count(INPUT, 40))
	fmt.Println(count(INPUT, 400000))
}
