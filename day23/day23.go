package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	reg := make(map[byte]int)
	// addition for part b
	//reg['a'] = 1

	for i := 0; i < len(instructions); i += 1 {
		instr := instructions[i]

		ins := instr[0:3]
		switch ins {
		case "hlf":
			reg[instr[4]] = reg[instr[4]] / 2
		case "tpl":
			reg[instr[4]] = reg[instr[4]] * 3
		case "inc":
			reg[instr[4]] = reg[instr[4]] + 1
		case "jmp":
			offset, e := strconv.Atoi(instr[4:])
			if e != nil {
				panic(e)
			}
			i += offset - 1
		case "jie":
			if reg[instr[4]]%2 == 0 {
				offset, e := strconv.Atoi(instr[7:])
				if e != nil {
					panic(e)
				}
				i += offset - 1
			}
		case "jio":
			if reg[instr[4]] == 1 {
				offset, e := strconv.Atoi(instr[7:])
				if e != nil {
					panic(e)
				}
				i += offset - 1
			}
		}
	}
	fmt.Println(reg)
}
