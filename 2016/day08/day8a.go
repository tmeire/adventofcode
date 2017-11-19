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

const (
	COLUMNS = 50
	ROWS    = 6
)

type Screen struct {
	p [][]bool
}

func New() *Screen {
	p := make([][]bool, ROWS)
	for i := 0; i < ROWS; i++ {
		p[i] = make([]bool, COLUMNS)
	}
	return &Screen{p}
}

func (s *Screen) RotateRow(row, n int) {
	for j := 0; j < n; j++ {
		tmp := s.p[row][COLUMNS-1]
		for i := COLUMNS - 2; i >= 0; i-- {
			s.p[row][i+1] = s.p[row][i]
		}
		s.p[row][0] = tmp
	}
}

func (s *Screen) RotateCol(col, n int) {
	for j := 0; j < n; j++ {
		tmp := s.p[ROWS-1][col]
		for i := ROWS - 2; i >= 0; i-- {
			s.p[i+1][col] = s.p[i][col]
		}
		s.p[0][col] = tmp
	}
}

func (s *Screen) Rect(x, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			s.p[i][j] = true
		}
	}
}

func (s *Screen) CountLit() int {
	count := 0
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			if s.p[i][j] {
				count++
			}
		}
	}
	return count
}
func (s *Screen) String() string {
	buf := bytes.NewBufferString("")
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			if j%5 == 0 && j > 0 {
				buf.WriteString(" ")
			}
			if s.p[i][j] {
				buf.WriteString("1")
			} else {
				buf.WriteString(" ")
			}
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func mustInt(a int, err error) int {
	if err != nil {
		panic(err)
	}
	return a
}

var REG1 = regexp.MustCompile(`rect ([0-9]+)x([0-9]+)`)
var REG2 = regexp.MustCompile(`rotate row y=([0-9]+) by ([0-9]+)`)
var REG3 = regexp.MustCompile(`rotate column x=([0-9]+) by ([0-9]+)`)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	screen := New()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmd := scanner.Text()

		r1 := REG1.FindStringSubmatch(cmd)
		if len(r1) > 0 {
			screen.Rect(mustInt(strconv.Atoi(r1[2])), mustInt(strconv.Atoi(r1[1])))
		} else {
			r2 := REG2.FindStringSubmatch(cmd)
			if len(r2) > 0 {
				screen.RotateRow(mustInt(strconv.Atoi(r2[1])), mustInt(strconv.Atoi(r2[2])))
			} else {
				r3 := REG3.FindStringSubmatch(cmd)
				if len(r3) > 0 {
					screen.RotateCol(mustInt(strconv.Atoi(r3[1])), mustInt(strconv.Atoi(r3[2])))
				}
			}
		}
	}
	fmt.Println("Lit pixels: ", screen.CountLit())
	fmt.Printf("%s\n", screen)
}
