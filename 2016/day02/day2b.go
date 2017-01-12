package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type button struct {
	up, down, left, right string
}

func (b button) get(d string) string {
	switch d {
	case "U":
		return b.up
	case "D":
		return b.down
	case "L":
		return b.left
	case "R":
		return b.right
	}
	return "0"
}

var keypad = map[string]button{
	"1": button{"0", "3", "0", "0"},
	"2": button{"0", "6", "0", "3"},
	"3": button{"1", "7", "2", "4"},
	"4": button{"0", "8", "3", "0"},
	"5": button{"0", "0", "0", "6"},
	"6": button{"2", "A", "5", "7"},
	"7": button{"3", "B", "6", "8"},
	"8": button{"4", "C", "7", "9"},
	"9": button{"0", "0", "8", "0"},
	"A": button{"6", "0", "0", "B"},
	"B": button{"7", "D", "A", "C"},
	"C": button{"8", "0", "B", "0"},
	"D": button{"B", "0", "0", "0"},
}

func main() {
	file, err := os.Open("input-a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b := "5"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, s := range scanner.Text() {
			bn := keypad[b].get(string(s))
			if bn != "0" {
				b = bn
			}
		}
		fmt.Printf(b)
	}
	fmt.Println()
}
