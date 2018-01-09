package day18

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
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

type CPU struct {
	id int

	reg     map[string]int
	program []instruction
	ip      int

	sends    int
	send     []int
	received []int

	sndchan chan int
	rcvchan chan int
}

func (cpu *CPU) load(program []instruction) {
	cpu.reg = make(map[string]int)

	for _, p := range program {
		p.preload(cpu)
	}
	cpu.reg["p"] = cpu.id

	cpu.program = program
	cpu.ip = 0
}

func (cpu *CPU) execute() bool {
	for cpu.ip >= 0 && cpu.ip < len(cpu.program) {
		//fmt.Println(cpu.id, cpu.ip, cpu.program[cpu.ip].needsinput(), len(cpu.received) == 0)
		if cpu.program[cpu.ip].needsinput() && len(cpu.received) == 0 {
			return false
		}
		cpu.ip += cpu.program[cpu.ip].apply(cpu)
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

func parse(s string) instruction {
	ss := CMD_REGEX.FindAllStringSubmatch(s, -1)
	switch ss[0][1] {
	case "snd":
		return SNDInstruction{SOI{ss[0][2]}}
	case "set":
		return SETInstruction{MOI{ss[0][2], ss[0][4]}}
	case "add":
		return ADDInstruction{MOI{ss[0][2], ss[0][4]}}
	case "mul":
		return MULInstruction{MOI{ss[0][2], ss[0][4]}}
	case "mod":
		return MODInstruction{MOI{ss[0][2], ss[0][4]}}
	case "rcv":
		return RCVInstruction{SOI{ss[0][2]}}
	case "jgz":
		return JGZInstruction{MOI{ss[0][2], ss[0][4]}}
	}
	panic("Unknown instruction " + s)
}

func load() []instruction {
	if len(os.Args) < 2 {
		panic("Must pass input file on commandline.")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cmds := []instruction{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmds = append(cmds, parse(scanner.Text()))
	}
	return cmds
}

func partA(cmds []instruction) {
	cpu := &CPU{}
	cpu.load(cmds)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		cpu.execute()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		reads := 0
		for a := range cpu.sendchan {
			if reads > 0 {
				<-cpu.readchan
			}
			cpu.readchan <- a
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Part A", cpu.send[len(cpu.send)-1])
}

func partB(cmds []instruction) {
	cpu1 := &CPU{id: 0}
	cpu1.load(cmds)

	cpu2 := &CPU{id: 1}
	cpu2.load(cmds)

	sends := 0

	var done1, done2 bool
	for sends == 0 || !((!done1 && len(cpu1.received) == 0) || (!done2 && (done1 || len(cpu1.received) == 0))) {
		done1 = cpu1.execute()

		cpu2.received = cpu1.send
		cpu1.send = nil

		done2 = cpu2.execute()

		cpu1.received = cpu2.send
		sends += len(cpu2.send)
		cpu2.send = nil
	}

	fmt.Println("Part B", cpu2.sends)
}

func Solve() {
	cmds := load()

	partA(cmds)
	partB(cmds)
}
