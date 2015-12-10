package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isHexNum(b byte) bool {
	// b is a number of character a,b,c,d,e or f
	return (48 <= b && b <= 57) || (97 <= b && b <= 102)
}

func cleaner(t string) string {
	t = strings.ToLower(t[1 : len(t)-1])
	s := make([]byte, 0, len(t))
	for i := 0; i < len(t); i += 1 {
		// if a normal character, just append
		if t[i] != '\\' {
			s = append(s, t[i])
			continue
		}

		// if last character, just append
		if i == len(t)-1 {
			s = append(s, t[i])
			continue
		}
		// if a backslash followed by a backslash, append backslash & skip character
		if t[i+1] == '\\' {
			s = append(s, '\\')
			i += 1
			continue
		}
		// if a backslash followed by a quote, append quote & skip character
		if t[i+1] == '"' {
			s = append(s, '"')
			i += 1
			continue
		}
		// if a backslash followed by xNN, append a dummy char & skip 3 characters
		if t[i+1] == 'x' && i+3 < len(t) && isHexNum(t[i+2]) && isHexNum(t[i+3]) {
			s = append(s, 'a')
			i += 3
			continue
		}
		s = append(s, '\\')
	}
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
	for scanner.Scan() {
		t := scanner.Text()

		total += len(t)
		clean += len(cleaner(t))
	}
	fmt.Println(total - clean)
}
