package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func containsVowels(s string) bool {
	vowels := 0
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels += 1
		}
	}
	return vowels >= 3
}

func containsRepeatedCharacters(s string) bool {
	for i := 1; i < len(s); i += 1 {
		if s[i] == s[i-1] {
			return true
		}
	}
	return false
}

func containsBadSubstring(s string) bool {
	return strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy")
}

func check(s string) bool {
	if !containsVowels(s) {
		return false
	}
	if !containsRepeatedCharacters(s) {
		return false
	}
	if containsBadSubstring(s) {
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
