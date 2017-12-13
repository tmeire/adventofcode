package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/blackskad/adventofcode/collection"
)

func parse(s string) (string, []string) {
	ss := strings.Split(s, " <-> ")

	return ss[0], strings.Split(ss[1], ", ")
}

func load(fname string) map[string][]string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pipes := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		o, dest := parse(scanner.Text())
		pipes[o] = dest
	}
	return pipes
}

func add(set collection.StringSet, pipes map[string][]string, origin string) {
	for _, d := range pipes[origin] {
		c := set.Contains(d)
		set.Put(d)
		if !c {
			add(set, pipes, d)
		}
	}
}

func Solve() {
	if len(os.Args) < 2 {
		fmt.Println("Expected the input file name as commandline argument.")
		os.Exit(1)
	}

	pipes := load(os.Args[1])

	set := collection.NewStringSet()
	set.Put("0")

	add(set, pipes, "0")

	fmt.Println("Part A:", len(set))

	sets := []collection.StringSet{set}

	for o := range pipes {
		found := false
		for _, s := range sets {
			if s.Contains(o) {
				found = true
				break
			}
		}
		if !found {
			newset := collection.NewStringSet()
			newset.Put(o)

			add(newset, pipes, o)
			sets = append(sets, newset)
		}
	}
	fmt.Println("Part B:", len(sets))
}
