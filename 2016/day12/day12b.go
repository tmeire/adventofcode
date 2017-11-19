package main

import "fmt"

func main() {
	cmds := load()

	cpu := &CPU{reg: make(map[string]int)}
	cpu.load(cmds)
	cpu.reg["c"] = 1
	cpu.execute()

	fmt.Println(cpu.reg["a"])
}
