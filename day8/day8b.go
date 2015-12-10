package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func cleaner(t string) string {
	s := make([]byte, 0, len(t)+2)
	s = append(s, '"')
	for i := 0; i < len(t); i += 1 {
		if t[i] == '\\' {
			s = append(s, '\\', '\\')
			continue
		}
		if t[i] == '"' {
			s = append(s, '\\', '"')
			continue
		}
		s = append(s, t[i])
	}
	s = append(s, '"')
	fmt.Println(string(s))
	return string(s)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	clean := 0

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		t := scanner.Text()

		total += len(t)
		tc := cleaner(t)
		clean += len(tc)
		fmt.Printf("%s - %s\n", t, tc)
		i += 1
	}
	fmt.Println(clean - total)
}
