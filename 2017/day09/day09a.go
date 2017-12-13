package day09

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var IGNORABLE = regexp.MustCompile(`!.`)
var GARBARGE = regexp.MustCompile(`<[^>]*>`)

func Solve() {
	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Remove all ignored characters from the stream
	bs = IGNORABLE.ReplaceAll(bs, []byte{})

	// Find all the garbage segments and count their length
	gbsize := 0
	gb := GARBARGE.FindAll(bs, -1)
	for _, g := range gb {
		gbsize += len(g) - 2
	}
	fmt.Println("Part B", gbsize)

	// Remove all the garbage from the stream
	bs = GARBARGE.ReplaceAll(bs, []byte{})

	// Keep track of the depth of the current group.
	// Add the depth to the sum each time we go deeper.
	sum := 0
	group := 0
	for _, b := range bs {
		switch b {
		case '{':
			group++
			sum += group
		case '}':
			group--
		}
	}
	fmt.Println("Part A", sum)
}

func Solve_v2() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Remove all ignored characters from the stream
	bs = IGNORABLE.ReplaceAll(bs, []byte{})

	gbsize := 0
	gb := GARBARGE.FindAll(bs, -1)
	for _, g := range gb {
		if len(g) > 1 {
			gbsize += len(g) - 2
		}
	}
	fmt.Println("Part B", gbsize)

	// Remove all the garbage from the stream
	bs = GARBARGE.ReplaceAll(bs, []byte{})

	sum := 0
	curg := 0
	for _, b := range bs {
		switch b {
		case '{':
			curg++
			sum += curg
		case '}':
			curg--
		}
	}
	fmt.Println("Part A", sum)
}
