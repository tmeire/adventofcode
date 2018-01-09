package collection

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type EdgeSet map[string][]string

func NewEdgeSetFromFile(fname string) EdgeSet {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	edges := make(EdgeSet)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), " <-> ")

		edges[ss[0]] = strings.Split(ss[1], ", ")
	}
	return edges
}

func (es EdgeSet) buildNodeSet(set StringSet, origin string) {
	for _, d := range es[origin] {
		if set.Put(d) {
			es.buildNodeSet(set, d)
		}
	}
}

func (es EdgeSet) NodeSets() []StringSet {
	sets := []StringSet{}

	for o := range es {
		found := false
		for _, s := range sets {
			if s.Contains(o) {
				found = true
				break
			}
		}
		if !found {
			newset := NewStringSet()
			newset.Put(o)

			es.buildNodeSet(newset, o)
			sets = append(sets, newset)
		}
	}
	return sets
}
