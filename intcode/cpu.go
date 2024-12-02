package intcode

const (
	opcodeADD = 1
	opcodeMUL = 2
	opcodeIN  = 3
	opcodeOUT = 4
	opcodeJIT = 5
	opcodeJIF = 6
	opcodeLT  = 7
	opcodeEQ  = 8
	opcodeRLB = 9
	opcodeEND = 99
)

const (
	position  = 0
	immediate = 1
)

// Program simulates a Program
type Program struct {
	mem     []int64
	ip      int64
	relbase int64

	Stdout chan int64
	Stdin  chan int64
	Done   chan bool
}

// NewProgram creates a new Program
func NewProgram(instructions []int64) *Program {
	return &Program{
		mem:     append(instructions, make([]int64, 4096)...),
		ip:      0,
		relbase: 0,

		Stdout: make(chan int64),
		Stdin:  make(chan int64),
		Done:   make(chan bool),
	}
}

func (p *Program) input() int64 {
	return <-p.Stdin
}

func pow10(n int64) int64 {
	var i int64 = 1
	for ; n > 0; n-- {
		i *= 10
	}
	return i
}

//(ins / 100) / (10 ^ (i-1)) % 10

func (p *Program) value(ins, i int64) int64 {
	mode := (ins / pow10(i+1)) % 10
	//println("VMODE", mode)
	switch mode {
	case 1:
		//println(p.ip + i)
		return p.mem[p.ip+i]
	case 2:
		// relative mode
		//println(p.relbase + p.mem[p.ip+i])
		return p.mem[p.relbase+p.mem[p.ip+i]]
	default:
		// position mode
		//println(p.mem[p.ip+i])
		return p.mem[p.mem[p.ip+i]]
	}
}

func (p *Program) store(ins, i, v int64) {
	mode := (ins / pow10(i+1)) % 10
	//println("SMODE", mode)
	switch mode {
	case 1:
		//println(p.ip + i)
		p.mem[p.ip+i] = v
	case 2:
		// relative mode
		//println(p.relbase + p.mem[p.ip+i])
		p.mem[p.relbase+p.mem[p.ip+i]] = v
	default:
		// position mode
		//println(p.mem[p.ip+i])
		p.mem[p.mem[p.ip+i]] = v
	}
}

func (p *Program) add(ins int64) int64 {
	//println(p.ip, ins, p.mem[p.ip+1], p.mem[p.ip+2], p.mem[p.ip+3])
	a1 := p.value(ins, 1)
	a2 := p.value(ins, 2)
	p.store(ins, 3, a1+a2)
	return 4
}

func (p *Program) mul(ins int64) int64 {
	//println(p.ip, ins, p.mem[p.ip+1], p.mem[p.ip+2], p.mem[p.ip+3])
	a1 := p.value(ins, 1)
	a2 := p.value(ins, 2)
	p.store(ins, 3, a1*a2)
	return 4
}

func (p *Program) in(ins int64) int64 {
	//println(p.ip, ins, p.mem[p.ip+1])
	p.store(ins, 1, p.input())
	return 2
}

func (p *Program) out(ins int64) int64 {
	//println(p.ip, ins, p.mem[p.ip+1])
	p.Stdout <- p.value(ins, 1)
	return 2
}

func (p *Program) jit(ins int64) int64 {
	//println(p.ip, ins, p.mem[p.ip+1], p.mem[p.ip+2])
	//println(p.value(ins, 1))
	if p.value(ins, 1) != 0 {
		p.ip = p.value(ins, 2)
		return 0
	}
	return 3
}

func (p *Program) jif(ins int64) int64 {
	//println(p.ip, ins, p.mem[p.ip+1], p.mem[p.ip+2])
	if p.value(ins, 1) == 0 {
		p.ip = p.value(ins, 2)
		return 0
	}
	return 3
}

func (p *Program) lt(ins int64) int64 {
	//println(p.ip, ins, p.mem[p.ip+1], p.mem[p.ip+2], p.mem[p.ip+3])
	if p.value(ins, 1) < p.value(ins, 2) {
		p.store(ins, 3, 1)
	} else {
		p.store(ins, 3, 0)
	}
	return 4
}

func (p *Program) eq(ins int64) int64 {
	//println(p.ip, ins, p.mem[p.ip+1], p.mem[p.ip+2], p.mem[p.ip+3])
	if p.value(ins, 1) == p.value(ins, 2) {
		p.store(ins, 3, 1)
	} else {
		p.store(ins, 3, 0)
	}
	return 4
}

func (p *Program) rlb(ins int64) int64 {
	//println(p.ip, ins, p.mem[p.ip+1])
	p.relbase += p.value(ins, 1)
	return 2
}

// Simulate runs the program
func (p *Program) Simulate() {
	quit := false
	for p.ip = 0; p.ip < int64(len(p.mem)) && !quit; {
		ins := p.mem[p.ip]
		opcode := ins % 100

		switch opcode {
		case opcodeADD:
			p.ip += p.add(ins)
		case opcodeMUL:
			p.ip += p.mul(ins)
		case opcodeIN:
			p.ip += p.in(ins)
		case opcodeOUT:
			p.ip += p.out(ins)
			//panic("OUTPUT SOMETHING :s")
		case opcodeJIT:
			p.ip += p.jit(ins)
		case opcodeJIF:
			p.ip += p.jif(ins)
		case opcodeLT:
			p.ip += p.lt(ins)
		case opcodeEQ:
			p.ip += p.eq(ins)
		case opcodeRLB:
			p.ip += p.rlb(ins)
		case opcodeEND:
			//fmt.Printf("%#v \n", p)
			quit = true
			p.Done <- true
		}
	}
}
