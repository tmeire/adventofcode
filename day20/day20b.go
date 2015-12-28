package main

import "fmt"

func main() {
	var s [50000000]int
	for i := 1; i <= 3603600; i += 1 {
		gifts := i * 11

		for j := 1; j <= 50; j += 1 {
			nr := i * j

			if nr >= len(s) {
				//fmt.Println("Array not big enough!")
				continue
			}
			s[nr] += gifts
		}
	}
	for idx, i := range s {
		if i >= 33100000 {
			fmt.Println(idx, i)
			return
		}
	}
}
