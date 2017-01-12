package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tls := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := scanner.Text()

		aba := make([]string, 0)
		bab := make([]string, 0)

		hypernet := false
		for i := 0; i < len(ip)-2; i++ {
			switch ip[i] {
			case '[':
				hypernet = true
			case ']':
				hypernet = false
			default:
				if ip[i] == ip[i+2] && ip[i] != ip[i+1] {
					if hypernet {
						bab = append(bab, ip[i:i+3])
					} else {
						aba = append(aba, ip[i:i+3])
					}
				}
			}
		}
		for _, x := range aba {
			cobab := make([]byte, 3)
			cobab[0] = x[1]
			cobab[1] = x[0]
			cobab[2] = x[1]
			fmt.Println(x, string(cobab), bab, contains(bab, string(cobab)))
			if contains(bab, string(cobab)) {
				tls++
				break
			}
		}
	}
	fmt.Println("TLS Capable: ", tls)
}
