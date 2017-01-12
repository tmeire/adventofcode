package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

		match := false
		hypernet := false
		hypernetMatch := false
		for i := 0; i < len(ip)-3; i++ {
			switch ip[i] {
			case '[':
				hypernet = true
			case ']':
				hypernet = false
			default:
				if ip[i] == ip[i+3] && ip[i+1] == ip[i+2] && ip[i] != ip[i+1] {
					if hypernet {
						hypernetMatch = true
					} else {
						match = true
					}
				}
			}
		}
		if match && !hypernetMatch {
			fmt.Println(ip, match, hypernetMatch)
			tls++
		}
	}
	fmt.Println("TLS Capable: ", tls)
}
