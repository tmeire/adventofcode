package main

import (
	"strings"

	"github.com/tmeire/adventofcode/io"
)

type object struct {
	name     string
	parent   *object
	children map[string]*object
}

func (obj *object) parents() []*object {
	if obj.parent == nil {
		return nil
	}
	return append(obj.parent.parents(), obj.parent)
}

func (obj *object) orbits() (int, int) {
	childCount := len(obj.children)
	childOrbits := len(obj.children)
	for _, child := range obj.children {
		cc, co := child.orbits()
		childCount += cc
		childOrbits += co
	}
	return childCount, childCount + childOrbits
}

func countOrbits(obj *object, depth int) int {
	if obj == nil {
		return 0
	}

	childOrbits := 0
	for _, child := range obj.children {
		childOrbits += countOrbits(child, depth+1)
	}
	return depth + childOrbits
}

func main() {

	orbits, err := io.ReadLinesFromFile("data.txt")
	if err != nil {
		panic(err)
	}
	/*orbits := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}*/

	objects := make(map[string]*object)

	for _, orbit := range orbits {
		orb := strings.Split(orbit, ")")

		obj1, ok := objects[orb[0]]
		if !ok {
			obj1 = &object{orb[0], nil, make(map[string]*object)}
			objects[orb[0]] = obj1
		}
		obj2, ok := objects[orb[1]]
		if !ok {
			obj2 = &object{orb[1], nil, make(map[string]*object)}
			objects[orb[1]] = obj2
		}

		obj1.children[orb[1]] = obj2
		obj2.parent = obj1
	}
	cc, orbitCount := objects["COM"].orbits()
	println(orbitCount - cc)
	println(countOrbits(objects["COM"], 0))

	p1 := objects["YOU"].parents()
	p2 := objects["SAN"].parents()

	common := 0
	for ; common < len(p1) && common < len(p2) && p1[common] == p2[common]; common++ {
		//
	}
	println((len(p2) + len(p1)) - 2*common)
}
