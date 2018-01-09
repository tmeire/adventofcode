package cpu

import (
	"bufio"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

/*
cpy x y copies x (either an integer or the value of a register) into register y.
inc x increases the value of register x by one.
dec x decreases the value of register x by one.
jnz x y jumps to an instruction y away (positive means forward; negative means backward), but only if x is not zero.
*/

type Instruction interface {
	preload(cpu *CPU)
	apply(cpu *CPU) int
	needsinput() bool
}

type SOI struct {
	o string
}

func (i SOI) preload(cpu *CPU) {
	if v, err := strconv.Atoi(i.o); err == nil {
		cpu.reg[i.o] = v
	}
}

func (i SOI) needsinput() bool {
	return false
}

type MOI struct {
	o1, o2 string
}

func (i MOI) preload(cpu *CPU) {
	if v, err := strconv.Atoi(i.o1); err == nil {
		cpu.reg[i.o1] = v
	}
	if v, err := strconv.Atoi(i.o2); err == nil {
		cpu.reg[i.o2] = v
	}
}

func (i MOI) needsinput() bool {
	return false
}

type SNDInstruction struct{ SOI }

func (i SNDInstruction) apply(cpu *CPU) int {
	cpu.snd(cpu.reg[i.o])
	return 1
}

type SETInstruction struct{ MOI }

func (i SETInstruction) apply(cpu *CPU) int {
	cpu.reg[i.o1] = cpu.reg[i.o2]
	return 1
}

type ADDInstruction struct{ MOI }

func (i ADDInstruction) apply(cpu *CPU) int {
	cpu.reg[i.o1] += cpu.reg[i.o2]
	return 1
}

type SUBInstruction struct{ MOI }

func (i SUBInstruction) apply(cpu *CPU) int {
	cpu.reg[i.o1] -= cpu.reg[i.o2]
	return 1
}

type MULInstruction struct{ MOI }

func (i MULInstruction) apply(cpu *CPU) int {
	cpu.reg[i.o1] *= cpu.reg[i.o2]
	return 1
}

type MODInstruction struct{ MOI }

func (i MODInstruction) apply(cpu *CPU) int {
	cpu.reg[i.o1] = cpu.reg[i.o1] % cpu.reg[i.o2]
	return 1
}

type RCVInstruction struct{ SOI }

func (i RCVInstruction) apply(cpu *CPU) int {
	cpu.reg[i.o] = cpu.rcv()
	return 1
}

func (i RCVInstruction) needsinput() bool {
	return true
}

type JGZInstruction struct{ MOI }

func (i JGZInstruction) apply(cpu *CPU) int {
	if cpu.reg[i.o1] > 0 {
		return cpu.reg[i.o2]
	}
	return 1
}

type JNZInstruction struct{ MOI }

func (i JNZInstruction) apply(cpu *CPU) int {
	if cpu.reg[i.o1] != 0 {
		return cpu.reg[i.o2]
	}
	return 1
}

type CPU struct {
	Debug  bool
	Counts map[string]int

	id int

	reg     map[string]int
	program []Instruction
	ip      int

	sends    int
	send     []int
	received []int
}

func (cpu *CPU) Load(program []Instruction) {
	cpu.Counts = make(map[string]int)
	cpu.reg = make(map[string]int)

	for _, p := range program {
		p.preload(cpu)
	}
	cpu.reg["p"] = cpu.id

	cpu.program = program
	cpu.ip = 0
}

func (cpu *CPU) Execute() bool {
	for cpu.ip >= 0 && cpu.ip < len(cpu.program) {
		ins := cpu.program[cpu.ip]

		if ins.needsinput() && len(cpu.received) == 0 {
			return false
		}
		if cpu.Debug {
			cpu.Counts[reflect.TypeOf(ins).Name()]++
		}

		cpu.ip += ins.apply(cpu)
	}
	return true
}

func (cpu *CPU) snd(a int) {
	cpu.send = append(cpu.send, a)
	cpu.sends++
}

func (cpu *CPU) rcv() int {
	tmp := cpu.received[0]
	cpu.received = cpu.received[1:]
	return tmp
}

var CMD_REGEX = regexp.MustCompile(`^([a-z]+) ([a-z]|[\-0-9]+)( ([a-z]|[\-0-9]+))?`)

func Parse(s string) Instruction {
	ss := CMD_REGEX.FindAllStringSubmatch(s, -1)
	switch ss[0][1] {
	case "snd":
		return SNDInstruction{SOI{ss[0][2]}}
	case "set":
		return SETInstruction{MOI{ss[0][2], ss[0][4]}}
	case "add":
		return ADDInstruction{MOI{ss[0][2], ss[0][4]}}
	case "sub":
		return SUBInstruction{MOI{ss[0][2], ss[0][4]}}
	case "mul":
		return MULInstruction{MOI{ss[0][2], ss[0][4]}}
	case "mod":
		return MODInstruction{MOI{ss[0][2], ss[0][4]}}
	case "rcv":
		return RCVInstruction{SOI{ss[0][2]}}
	case "jgz":
		return JGZInstruction{MOI{ss[0][2], ss[0][4]}}
	case "jnz":
		return JNZInstruction{MOI{ss[0][2], ss[0][4]}}
	}
	panic("Unknown instruction " + s)
}

func Load(fname string) []Instruction {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cmds := []Instruction{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmds = append(cmds, Parse(scanner.Text()))
	}
	return cmds
}
