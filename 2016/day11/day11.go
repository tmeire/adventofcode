package main

import (
	"fmt"
	"sync"

	"github.com/cnf/structhash"
)

type OType int

const (
	CHIP OType = iota
	GEN
)

type Element int64

const (
	NOOP       Element = 0
	THULIUM    Element = 1 << 0
	RUTHENIUM  Element = 1 << 1
	COBALT     Element = 1 << 2
	POLONIUM   Element = 1 << 3
	PROMETHIUM Element = 1 << 4
	ELERIUM    Element = 1 << 5
	DILITHIUM  Element = 1 << 6
)

var ELEMENTS = []Element{THULIUM, RUTHENIUM, COBALT, POLONIUM, PROMETHIUM}

type object struct {
	otype   OType
	element Element
}

//---------------------------------------------------------------------

type Direction int

const (
	UP   Direction = 1
	DOWN Direction = -1
)

type move struct {
	dir    Direction
	e1, e2 object
}

//---------------------------------------------------------------------

type floor struct {
	Gens, Chips int64
}

func (f floor) empty() bool {
	return f.Gens == 0 && f.Chips == 0
}

func (f floor) hash() int64 {
	return f.Chips | (f.Gens << 10)
}

func (f floor) burned() bool {
	if f.Gens == 0 {
		// Chips can't get burned if there are no generators on the floor
		return false
	}

	return (f.Chips &^ f.Gens) != 0
}

func (f *floor) add(o object) {
	switch o.otype {
	case CHIP:
		f.Chips = f.Chips | int64(o.element)
	case GEN:
		f.Gens = f.Gens | int64(o.element)
	}
}

func (f *floor) remove(o object) {
	switch o.otype {
	case CHIP:
		f.Chips = f.Chips &^ int64(o.element)
	case GEN:
		f.Gens = f.Gens &^ int64(o.element)
	}
}

func (f floor) containsChip(e Element) bool {
	return f.Chips&int64(e) == int64(e)
}

func (f floor) containsGenerator(e Element) bool {
	return f.Gens&int64(e) == int64(e)
}

func (f floor) items() int {
	c := 0
	for _, e := range ELEMENTS {
		if f.containsChip(e) {
			c++
		}
		if f.containsGenerator(e) {
			c++
		}
	}
	return c
}

//---------------------------------------------------------------------

type facility struct {
	moves    int `hash:"-"`
	Elevator int

	Floors []floor

	next *facility `hash:"-"`
}

func (f facility) burned() bool {
	for _, fl := range f.Floors {
		if fl.burned() {
			return true
		}
	}
	return false
}

func (f facility) completed() bool {
	return f.Floors[0].empty() && f.Floors[1].empty() && f.Floors[2].empty() && !f.Floors[3].empty()
}

func (f facility) nonEmptyLower() bool {
	for i := 0; i < f.Elevator; i++ {
		if !f.Floors[i].empty() {
			return true
		}
	}
	return false
}

func (f facility) nextMoves() []move {
	ms := []move{}

	// Combine all items on the current floor in groups of 2 to move them up
	if f.Elevator < 3 {
		// chips with gens
		for _, e1 := range ELEMENTS {
			if f.Floors[f.Elevator].containsChip(e1) {
				for _, e2 := range ELEMENTS {
					if f.Floors[f.Elevator].containsGenerator(e2) {
						ms = append(ms, move{dir: UP, e1: object{CHIP, e1}, e2: object{GEN, e2}})
					}
				}
			}
		}
		// chips with chips
		for i, e1 := range ELEMENTS {
			if f.Floors[f.Elevator].containsChip(e1) {
				for j, e2 := range ELEMENTS {
					if j > i && f.Floors[f.Elevator].containsChip(e2) {
						ms = append(ms, move{dir: UP, e1: object{CHIP, e1}, e2: object{CHIP, e2}})
					}
				}
			}
		}
		// gens with gens
		for i, e1 := range ELEMENTS {
			if f.Floors[f.Elevator].containsGenerator(e1) {
				for j, e2 := range ELEMENTS {
					if j > i && f.Floors[f.Elevator].containsGenerator(e2) {
						ms = append(ms, move{dir: UP, e1: object{GEN, e1}, e2: object{GEN, e2}})
					}
				}
			}
		}
		// Try to move up with a single element
		for _, e1 := range ELEMENTS {
			if f.Floors[f.Elevator].containsChip(e1) {
				ms = append(ms, move{dir: UP, e1: object{CHIP, e1}})
			}
			if f.Floors[f.Elevator].containsGenerator(e1) {
				ms = append(ms, move{dir: UP, e1: object{GEN, e1}})
			}
		}
	}

	// Try to move down with each item individually
	if f.nonEmptyLower() && f.Elevator > 0 {
		// Try to move down with a single element
		for _, e1 := range ELEMENTS {
			if f.Floors[f.Elevator].containsChip(e1) {
				ms = append(ms, move{dir: DOWN, e1: object{CHIP, e1}})
			}
			if f.Floors[f.Elevator].containsGenerator(e1) {
				ms = append(ms, move{dir: DOWN, e1: object{GEN, e1}})
			}
		}
	}
	return ms
}

func (f *facility) clone() *facility {
	c := *f
	c.Floors = append([]floor(nil), f.Floors...)
	return &c
}

func (f *facility) apply(m move) *facility {
	switch m.dir {
	case UP:
		f.Floors[f.Elevator].remove(m.e1)
		f.Floors[f.Elevator+1].add(m.e1)
		if m.e2.element != NOOP {
			f.Floors[f.Elevator].remove(m.e2)
			f.Floors[f.Elevator+1].add(m.e2)
		}
		f.Elevator += 1
	case DOWN:
		f.Floors[f.Elevator].remove(m.e1)
		f.Floors[f.Elevator-1].add(m.e1)
		f.Elevator -= 1
	}
	f.moves += 1
	return f
}

func (f facility) score() int {
	// It minimally takes (n-1)*2-1 moves to move n items one floor up
	estimate := 0
	estimate += (f.Floors[0].items()) / 2 * 12
	estimate += (f.Floors[1].items()) / 2 * 8
	estimate += (f.Floors[2].items()) / 2 * 4
	return f.moves + estimate
}

func (f *facility) hash() string {
	return fmt.Sprintf("%x", structhash.Md5(f, 1))
}

func (f *facility) print() {
	e2c := func(floor int) string {
		if floor == f.Elevator {
			return ">"
		}
		return " "
	}

	fmt.Println("------")
	fmt.Printf("%s%015b\n", e2c(3), f.Floors[3].hash())
	fmt.Printf("%s%015b\n", e2c(2), f.Floors[2].hash())
	fmt.Printf("%s%015b\n", e2c(1), f.Floors[1].hash())
	fmt.Printf("%s%015b\n", e2c(0), f.Floors[0].hash())
	fmt.Println("------")
}

func (f *facility) queue(fNext *facility) {
	s := fNext.score()

	p := f

	for p.next != nil && p.next.score() <= s {
		p = p.next
	}
	fNext.next = p.next
	p.next = fNext
}

//---------------------------------------------------------------------

func collect(f *facility) int {
	visited := sync.Map{}

	cpu := 8

	// Make result buffered with room for each goroutine to return a response, to make sure they don't block
	result := make(chan int, cpu)
	// Channel to signal the goroutines that they can stop
	done := make(chan struct{})
	// A queue channel containing remaining work
	queue := make(chan *facility, 100000000)

	wg := sync.WaitGroup{}
	for i := 0; i < cpu; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				case f := <-queue:
					//f.print()
					h := f.hash()

					// if we already visited f in the meantime, skip it
					if _, ok := visited.Load(h); ok {
						continue
					}

					for _, m := range f.nextMoves() {
						fNew := f.clone().apply(m)
						if fNew.completed() {
							result <- fNew.moves
							return
						}
						if fNew.burned() {
							continue
						}

						// if we already visited fNew, don't add it to the queue
						if _, ok := visited.Load(fNew.hash()); ok {
							continue
						}

						queue <- fNew
						//f.queue(fNew)
					}
					visited.Store(h, struct{}{})
				}
			}
		}()
	}
	// Kick of the processing in the goroutines
	queue <- f

	// Wait for the first one to finish
	r := <-result

	// Signal the others to quit
	close(done)

	// Wait for the others to quit
	wg.Wait()

	// Close the channels
	close(queue)
	close(result)

	return r
}

func main() {
	start := &facility{
		moves:    0,
		Elevator: 0,
		Floors: []floor{
			floor{}, floor{}, floor{}, floor{},
		},
		next: nil,
	}

	// Try the last steps, should print 5
	start.Elevator = 2
	start.Floors[2].add(object{CHIP, THULIUM})
	start.Floors[2].add(object{CHIP, RUTHENIUM})
	start.Floors[2].add(object{GEN, THULIUM})
	start.Floors[2].add(object{GEN, RUTHENIUM})

	fmt.Println(collect(start))

	// Try the complete example, should print 11
	start = &facility{
		moves:    0,
		Elevator: 0,
		Floors: []floor{
			floor{}, floor{}, floor{}, floor{},
		},
		next: nil,
	}
	start.Floors[0].add(object{CHIP, THULIUM})
	start.Floors[0].add(object{CHIP, RUTHENIUM})
	start.Floors[1].add(object{GEN, THULIUM})
	start.Floors[2].add(object{GEN, RUTHENIUM})

	fmt.Println(collect(start))

	// Try the actual puzzle
	start = &facility{
		moves:    0,
		Elevator: 0,
		Floors: []floor{
			floor{}, floor{}, floor{}, floor{},
		},
		next: nil,
	}

	start.Floors[0].add(object{CHIP, THULIUM})
	start.Floors[0].add(object{CHIP, RUTHENIUM})
	start.Floors[0].add(object{CHIP, COBALT})
	start.Floors[0].add(object{GEN, THULIUM})
	start.Floors[0].add(object{GEN, RUTHENIUM})
	start.Floors[0].add(object{GEN, COBALT})
	start.Floors[0].add(object{GEN, POLONIUM})
	start.Floors[0].add(object{GEN, PROMETHIUM})

	start.Floors[1].add(object{CHIP, POLONIUM})
	start.Floors[1].add(object{CHIP, PROMETHIUM})

	fmt.Println(collect(start))

	// Try the actual puzzle, part b
	ELEMENTS = append(ELEMENTS, ELERIUM, DILITHIUM)

	start = &facility{
		moves:    0,
		Elevator: 0,
		Floors: []floor{
			floor{}, floor{}, floor{}, floor{},
		},
		next: nil,
	}
	start.Floors[0].add(object{CHIP, THULIUM})
	start.Floors[0].add(object{CHIP, RUTHENIUM})
	start.Floors[0].add(object{CHIP, COBALT})
	start.Floors[0].add(object{GEN, THULIUM})
	start.Floors[0].add(object{GEN, RUTHENIUM})
	start.Floors[0].add(object{GEN, COBALT})
	start.Floors[0].add(object{GEN, POLONIUM})
	start.Floors[0].add(object{GEN, PROMETHIUM})

	start.Floors[1].add(object{CHIP, POLONIUM})
	start.Floors[1].add(object{CHIP, PROMETHIUM})

	start.Floors[0].add(object{CHIP, ELERIUM})
	start.Floors[0].add(object{CHIP, DILITHIUM})
	start.Floors[0].add(object{GEN, ELERIUM})
	start.Floors[0].add(object{GEN, DILITHIUM})

	fmt.Println(collect(start))
}
