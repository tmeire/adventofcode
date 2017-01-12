package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	possible := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		a1, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b1, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		c1, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		a2, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b2, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		c2, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		a3, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b3, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		c3, _ := strconv.Atoi(scanner.Text())

		if check(a1, a2, a3) {
			possible++
		}
		if check(b1, b2, b3) {
			possible++
		}
		if check(c1, c2, c3) {
			possible++
		}
	}
	fmt.Println("Possible triangles: ", possible)
}
