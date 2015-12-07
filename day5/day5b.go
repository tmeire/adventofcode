package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func containsRepeatedCharacterGroups(s string) bool {
	for i := 0; i < len(s)-3; i += 1 {
		for j := i + 2; j < len(s)-1; j += 1 {
			if s[i:i+2] == s[j:j+2] {
				return true
			}
		}
	}
	return false
}

func containsRepeatedCharactersWithSkip(s string) bool {
	for i := 2; i < len(s); i += 1 {
		if s[i] == s[i-2] {
			return true
		}
	}
	return false
}

func check(s string) bool {
	if !containsRepeatedCharacterGroups(s) {
		return false
	}
	if !containsRepeatedCharactersWithSkip(s) {
		return false
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nice := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if check(scanner.Text()) {
			nice += 1
		}
	}
	fmt.Printf("Nice lines: %d\n", nice)
}
