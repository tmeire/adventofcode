package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var MAXDEPTH = 10

func atoi(s string) uint16 {
	i, _ := strconv.ParseUint(s, 10, 16)
	return uint16(i)
}

type Value interface {
	compute() uint16
}

type computed struct {
	value    uint16
	computed bool
}

type Root struct {
	computed
	input uint16
}

func (r *Root) compute() uint16 {
	if !r.computed.computed {
		r.computed.value = r.input
		r.computed.computed = true
	}
	return r.computed.value
}

type Not struct {
	computed
	input Value
}

func (n *Not) compute() uint16 {
	if !n.computed.computed {
		n.computed.value = ^n.input.compute()
		n.computed.computed = true
	}
	return n.computed.value
}

type And struct {
	computed
	inputA Value
	inputB Value
}

func (a *And) compute() uint16 {
	if !a.computed.computed {
		a.computed.value = a.inputA.compute() & a.inputB.compute()
		a.computed.computed = true
	}
	return a.computed.value
}

type Or struct {
	computed
	inputA Value
	inputB Value
}

func (o *Or) compute() uint16 {
	if !o.computed.computed {
		o.computed.value = o.inputA.compute() | o.inputB.compute()
		o.computed.computed = true
	}
	return o.computed.value
}

type LShift struct {
	computed
	input Value
	shift uint16
}

func (l *LShift) compute() uint16 {
	if !l.computed.computed {
		l.computed.value = l.input.compute() << l.shift
		l.computed.computed = true
	}
	return l.computed.value
}

type RShift struct {
	computed
	input Value
	shift uint16
}

func (r *RShift) compute() uint16 {
	if !r.computed.computed {
		r.computed.value = r.input.compute() >> r.shift
		r.computed.computed = true
	}
	return r.computed.value
}

type Definition struct {
	parts []string
}

var defs map[string]*Definition
var vals map[string]Value

func getCachedValueForInput(name string) Value {
	if vals[name] != nil {
		return vals[name]
	}
	v := getValueForInput(name)
	vals[name] = v
	return v
}

func getValueForInput(name string) Value {
	// if the name can be parsed as a number, return it as a root value
	v, err := strconv.ParseUint(name, 10, 16)
	if err == nil {
		return &Root{computed{}, uint16(v)}
	}

	// find the definition for the name
	d := defs[name]

	// uncomment this for part b
	//if name == "b" {
	//	return &Root{computed{}, uint16(46065)}
	//}

	// Create a new value based on the operation
	switch d.parts[1] {
	case "": // basically a passthrough / assignment
		return getCachedValueForInput(d.parts[2])
	case "NOT":
		return &Not{computed{}, getCachedValueForInput(d.parts[2])}
	case "AND":
		return &And{computed{}, getCachedValueForInput(d.parts[0]), getCachedValueForInput(d.parts[2])}
	case "OR":
		return &Or{computed{}, getCachedValueForInput(d.parts[0]), getCachedValueForInput(d.parts[2])}
	case "LSHIFT":
		return &LShift{computed{}, getCachedValueForInput(d.parts[0]), atoi(d.parts[2])}
	case "RSHIFT":
		return &RShift{computed{}, getCachedValueForInput(d.parts[0]), atoi(d.parts[2])}
	}
	return nil
}

var regex = regexp.MustCompile("([a-z0-9]*?)?[ ]?([A-Z]*)?[ ]?([a-z0-9]+) -> ([a-z]*)")

func main() {
	defs = make(map[string]*Definition)
	vals = make(map[string]Value)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := regex.FindStringSubmatch(scanner.Text())
		if m == nil {
			continue
		}
		d := Definition{m[1:5]}
		defs[d.parts[3]] = &d
	}
	fmt.Printf("Value of a: %d\n", getCachedValueForInput("a").compute())
}
