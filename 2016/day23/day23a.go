package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

/*
cpy x y copies x (either an integer or the value of a register) into register y.
inc x increases the value of register x by one.
dec x decreases the value of register x by one.
jnz x y jumps to an instruction y away (positive means forward; negative means backward), but only if x is not zero.
*/

type instruction interface {
	preload(cpu *CPU)
	apply(cpu *CPU) int
	toggle() instruction
}

type CPYInstruction struct {
	o1, o2 string
}

func (i CPYInstruction) preload(cpu *CPU) {
	if v, err := strconv.Atoi(i.o1); err == nil {
		cpu.reg[i.o1] = v
	}
	if v, err := strconv.Atoi(i.o2); err == nil {
		cpu.reg[i.o2] = v
	}
}

func (i CPYInstruction) apply(cpu *CPU) int {
	if cpu.isreg(i.o2) {
		cpu.reg[i.o2] = cpu.reg[i.o1]
	}
	return 1
}

func (i CPYInstruction) toggle() instruction {
	return JNZInstruction{i.o1, i.o2}
}

type INCInstruction struct {
	o string
}

func (i INCInstruction) preload(cpu *CPU) {
	if v, err := strconv.Atoi(i.o); err == nil {
		cpu.reg[i.o] = v
	}
}

func (i INCInstruction) apply(cpu *CPU) int {
	if cpu.isreg(i.o) {
		cpu.reg[i.o] += 1
	}
	return 1
}

func (i INCInstruction) toggle() instruction {
	return DECInstruction{i.o}
}

type DECInstruction struct {
	o string
}

func (i DECInstruction) preload(cpu *CPU) {
	if v, err := strconv.Atoi(i.o); err == nil {
		cpu.reg[i.o] = v
	}
}

func (i DECInstruction) apply(cpu *CPU) int {
	if cpu.isreg(i.o) {
		cpu.reg[i.o] -= 1
	}
	return 1
}

func (i DECInstruction) toggle() instruction {
	return INCInstruction{i.o}
}

type JNZInstruction struct {
	o1, o2 string
}

func (i JNZInstruction) preload(cpu *CPU) {
	if v, err := strconv.Atoi(i.o1); err == nil {
		cpu.reg[i.o1] = v
	}
	if v, err := strconv.Atoi(i.o2); err == nil {
		cpu.reg[i.o2] = v
	}
}

func (i JNZInstruction) apply(cpu *CPU) int {
	if cpu.reg[i.o1] != 0 {
		return cpu.reg[i.o2]
	}
	return 1
}

func (i JNZInstruction) toggle() instruction {
	return CPYInstruction{i.o1, i.o2}
}

type TGLInstruction struct {
	o string
}

func (i TGLInstruction) preload(cpu *CPU) {
	if v, err := strconv.Atoi(i.o); err == nil {
		cpu.reg[i.o] = v
	}
}

func (i TGLInstruction) apply(cpu *CPU) int {
	idx := cpu.ip + cpu.reg[i.o]
	if idx >= 0 && idx < len(cpu.program) {
		cpu.program[idx] = cpu.program[idx].toggle()
	}
	return 1
}

func (i TGLInstruction) toggle() instruction {
	return INCInstruction{i.o}
}

type CPU struct {
	reg     map[string]int
	program []instruction
	ip      int
}

func (cpu *CPU) isreg(s string) bool {
	return s == "a" || s == "b" || s == "c" || s == "d"
}

func (cpu *CPU) load(program []instruction) {
	for _, p := range program {
		p.preload(cpu)
	}

	cpu.program = program
	cpu.ip = 0

	cpu.reg["a"] = 0
	cpu.reg["b"] = 0
	cpu.reg["c"] = 0
	cpu.reg["d"] = 0
}

func (cpu *CPU) execute() {
	for cpu.ip = 0; cpu.ip < len(cpu.program); {
		cpu.ip += cpu.program[cpu.ip].apply(cpu)

		/*for _, ins := range cpu.program {
			fmt.Printf("%#v\n", ins)
		}
		fmt.Println("---")*/
	}
}

var CMD_REGEX = regexp.MustCompile(`([a-z]+) ([a-d]|[\-0-9]+)( ([a-d]|[\-0-9]+))?`)

func parse(s string) instruction {
	ss := CMD_REGEX.FindAllStringSubmatch(s, -1)
	switch ss[0][1] {
	case "cpy":
		return CPYInstruction{ss[0][2], ss[0][4]}
	case "inc":
		return INCInstruction{ss[0][2]}
	case "dec":
		return DECInstruction{ss[0][2]}
	case "jnz":
		return JNZInstruction{ss[0][2], ss[0][4]}
	case "tgl":
		return TGLInstruction{ss[0][2]}
	}
	panic("Unknown instruction " + s)
}

func load() []instruction {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cmds := []instruction{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmds = append(cmds, parse(scanner.Text()))
	}
	return cmds
}

func main() {
	cmds := load()

	cpu := &CPU{reg: make(map[string]int)}
	cpu.load(cmds)
	cpu.reg["a"] = 7
	cpu.reg["a"] = 12
	cpu.execute()

	fmt.Println(cpu.reg["a"])
}
