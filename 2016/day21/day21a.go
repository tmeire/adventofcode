package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var SWAP_POS_REGEX = regexp.MustCompile(`swap position ([0-9]) with position ([0-9])`)
var SWAP_LETTER_REGEX = regexp.MustCompile(`swap letter ([a-z]) with letter ([a-z])`)
var ROTATE_REGEX = regexp.MustCompile(`rotate (left|right) ([0-9]) step[s]?`)
var ROTATE_POS_REGEX = regexp.MustCompile(`rotate based on position of letter ([a-z])`)
var REVERSE_REGEX = regexp.MustCompile(`reverse positions ([0-9]) through ([0-9])`)
var MOVE_REGEX = regexp.MustCompile(`move position ([0-9]) to position ([0-9])`)

func stringToDirection(s []byte) RotationDirection {
	if string(s) == "left" {
		return LEFT
	}
	return RIGHT
}

func parse(s string) Operation {
	b := []byte(s)

	var ss [][]byte

	ss = SWAP_POS_REGEX.FindSubmatch(b)
	if len(ss) != 0 {
		x, _ := strconv.Atoi(string(ss[1]))
		y, _ := strconv.Atoi(string(ss[2]))
		return SwapPos{x, y}
	}

	ss = SWAP_LETTER_REGEX.FindSubmatch(b)
	if len(ss) != 0 {
		return SwapLetter{ss[1], ss[2]}
	}

	ss = ROTATE_REGEX.FindSubmatch(b)
	if len(ss) != 0 {
		x, _ := strconv.Atoi(string(ss[2]))
		return RotateX{stringToDirection(ss[1]), x}
	}

	ss = ROTATE_POS_REGEX.FindSubmatch(b)
	if len(ss) != 0 {
		return RotateLetter{ss[1]}
	}

	ss = REVERSE_REGEX.FindSubmatch(b)
	if len(ss) != 0 {
		x, _ := strconv.Atoi(string(ss[1]))
		y, _ := strconv.Atoi(string(ss[2]))
		return Reverse{x, y}
	}

	ss = MOVE_REGEX.FindSubmatch(b)
	if len(ss) != 0 {
		x, _ := strconv.Atoi(string(ss[1]))
		y, _ := strconv.Atoi(string(ss[2]))
		return Move{x, y}
	}

	fmt.Println("UNKNOWN COMMAND ", s)
	return nil
}

func load() []Operation {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ops := []Operation{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ops = append(ops, parse(scanner.Text()))
	}
	return ops
}

type RotationDirection bool

const (
	LEFT  RotationDirection = true
	RIGHT RotationDirection = false
)

func rotate(s []byte, dir RotationDirection, x int) []byte {
	var splitIDX int

	switch dir {
	case LEFT:
		// rotate left
		splitIDX = x
	case RIGHT:
		// rotate right
		splitIDX = len(s) - x
	}

	r := make([]byte, 0, len(s))
	r = append(r, s[splitIDX:]...)
	r = append(r, s[0:splitIDX]...)
	return r
}

type Operation interface {
	execute([]byte) []byte
	undo([]byte) []byte
}

type SwapPos struct {
	x, y int
}

func (k SwapPos) swap(s []byte, x, y int) []byte {
	s[x], s[y] = s[y], s[x]
	return s
}

func (k SwapPos) execute(s []byte) []byte {
	return k.swap(s, k.x, k.y)
}

func (k SwapPos) undo(s []byte) []byte {
	return k.swap(s, k.y, k.x)
}

type SwapLetter struct {
	x, y []byte
}

func (k SwapLetter) swap(s []byte, x, y []byte) []byte {
	r := bytes.Replace(s, x, []byte{'x'}, -1)
	r = bytes.Replace(r, y, x, -1)
	return bytes.Replace(r, []byte{'x'}, y, -1)
}

func (k SwapLetter) execute(s []byte) []byte {
	return k.swap(s, k.x, k.y)
}

func (k SwapLetter) undo(s []byte) []byte {
	return k.swap(s, k.y, k.x)
}

type RotateX struct {
	dir RotationDirection
	x   int
}

func (k RotateX) execute(s []byte) []byte {
	return rotate(s, k.dir, k.x)
}

func (k RotateX) undo(s []byte) []byte {
	return rotate(s, !k.dir, k.x)
}

type RotateLetter struct {
	x []byte
}

func (k RotateLetter) execute(s []byte) []byte {
	idx := bytes.Index(s, k.x)

	opt := 0
	if idx >= 4 {
		opt = 1
	}
	rot := (1 + idx + opt) % len(s)
	return rotate(s, RIGHT, rot)
}

func (k RotateLetter) undo(s []byte) []byte {
	idx := bytes.Index(s, k.x)

	var nidx int
	if idx == 0 {
		nidx = (2*len(s) + idx - 2) / 2
	} else if idx%2 == 0 {
		nidx = (len(s) + idx - 2) / 2
	} else {
		nidx = (idx - 1) / 2
	}

	return rotate(s, RIGHT, (nidx-idx+len(s))%len(s))
}

type Reverse struct {
	x, y int
}

func (k Reverse) execute(s []byte) []byte {
	rev := make([]byte, k.y-k.x+1)
	for i := k.y; i >= k.x; i-- {
		rev[k.y-i] = s[i]
	}

	r := make([]byte, 0, len(s))

	// If there's a part before the reversal, add it
	if k.x > 0 {
		r = append(r, s[:k.x]...)
	}

	// Add the reversal
	r = append(r, rev...)

	// If there's a part after the reversal, add it
	if k.y < len(s)-1 {
		r = append(r, s[k.y+1:]...)
	}
	return r
}

func (k Reverse) undo(s []byte) []byte {
	return k.execute(s)
}

type Move struct {
	x, y int
}

func (k Move) move(s []byte, x, y int) []byte {
	r := make([]byte, 0, len(s))

	if x < y {
		// The part before x
		r = append(r, s[0:x]...)
		// The part between x and y
		r = append(r, s[x+1:y+1]...)
		// Insert x at y
		r = append(r, s[x])
		// The remaining part after y
		r = append(r, s[y+1:]...)
	} else {
		// The part before y
		r = append(r, s[0:y]...)
		// Insert x at y
		r = append(r, s[x])
		// The part between y and x
		r = append(r, s[y:x]...)
		// The remaining part after x
		r = append(r, s[x+1:]...)
	}
	return r
}

func (k Move) execute(s []byte) []byte {
	return k.move(s, k.x, k.y)
}

func (k Move) undo(s []byte) []byte {
	return k.move(s, k.y, k.x)
}

func main() {
	s := []byte("abcdefgh")
	oss := string(s)

	ops := load()

	for _, o := range ops {
		//fmt.Printf("%d %#v %s", i, o, s)
		s = o.execute(s)
	}
	fmt.Println(string(s))
	fmt.Println("----------------------")

	s = []byte("fbgdceah")
	for i := len(ops) - 1; i >= 0; i-- {
		s = ops[i].undo(s)
	}
	fmt.Println(string(s))
	fmt.Println(oss == string(s))
}
