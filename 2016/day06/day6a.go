package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func biggest(data map[rune]int) rune {
	bigr := 'a'
	bigc := 0
	for k, v := range data {
		if v > bigc {
			bigc = v
			bigr = k
		}
	}
	return bigr
}

func smallest(data map[rune]int) rune {
	bigr := 'a'
	bigc := 10000
	for k, v := range data {
		if v < bigc {
			bigc = v
			bigr = k
		}
	}
	return bigr
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make(map[int]map[rune]int)
	for i := 0; i < 8; i++ {
		data[i] = make(map[rune]int)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for i, r := range scanner.Text() {
			data[i][r]++
		}
	}

	word := make([]rune, 8)
	for i, d := range data {
		// word[i] = biggest(d)
		word[i] = smallest(d)
	}
	fmt.Println(string(word))
}
