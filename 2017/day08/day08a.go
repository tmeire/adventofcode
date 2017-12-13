package day08

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type condition struct {
	reg string
	op  string
	val int
}

func (c *condition) evaluate(cpu *CPU) bool {
	switch c.op {
	case "==":
		return cpu.reg[c.reg] == c.val
	case "!=":
		return cpu.reg[c.reg] != c.val
	case "<":
		return cpu.reg[c.reg] < c.val
	case ">":
		return cpu.reg[c.reg] > c.val
	case "<=":
		return cpu.reg[c.reg] <= c.val
	case ">=":
		return cpu.reg[c.reg] >= c.val
	}
	panic("Unknown condition " + c.op)
}

type instruction struct {
	reg string
	op  string
	val int
	c   *condition
}

func (i *instruction) execute(cpu *CPU) {
	if !i.c.evaluate(cpu) {
		return
	}

	switch i.op {
	case "inc":
		cpu.add(i.reg, i.val)
	case "dec":
		cpu.add(i.reg, -i.val)
	}
}

type CPU struct {
	reg map[string]int

	highValue int
}

func (cpu *CPU) add(reg string, val int) {
	cpu.reg[reg] += val
	if cpu.reg[reg] > cpu.highValue {
		cpu.highValue = cpu.reg[reg]
	}
}

func (cpu *CPU) max() int {
	max := 0
	for _, v := range cpu.reg {
		if v > max {
			max = v
		}
	}
	return max
}

func (cpu *CPU) execute(program []*instruction) {
	for _, i := range program {
		i.execute(cpu)
	}
}

func load(fname string) []*instruction {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cmds := []*instruction{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := instruction{c: &condition{}}

		fmt.Sscanf(scanner.Text(), "%s %s %d if %s %s %d", &(i.reg), &(i.op), &(i.val), &(i.c.reg), &(i.c.op), &(i.c.val))

		cmds = append(cmds, &i)
	}
	return cmds
}

func Solve() {
	if len(os.Args) < 2 {
		fmt.Println("Expected the input file name as commandline argument.")
		os.Exit(1)
	}

	cmds := load(os.Args[1])

	cpu := &CPU{}
	cpu.execute(cmds)

	fmt.Println("Part A", cpu.max())
	fmt.Println("Part B", cpu.highValue)
}
