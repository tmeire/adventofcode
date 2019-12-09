package main

import "strconv"

func main() {
	found := 0
	for x := 168630; x < 718098; x++ {
		s := []byte(strconv.Itoa(x))
		d := false
		u := true
		for i := 1; i < len(s); i++ {
			// Check if numbers are incremental
			if s[i] < s[i-1] {
				u = false
			}
			if s[i] == s[i-1] {
				d = true
			}
		}
		if d && u {
			println(x)
			found++
		}
	}
	println(found)

	found = 0
	for x := 168630; x < 718098; x++ {
		s := []byte(strconv.Itoa(x))
		u := true
		n := make(map[byte]int)
		n[s[0]]++
		for i := 1; i < len(s); i++ {
			// Check if numbers are incremental
			if s[i] < s[i-1] {
				u = false
			}
			n[s[i]]++
		}
		d := false
		for _, v := range n {
			if v == 2 {
				d = true
			}
		}
		if d && u {
			println(x)
			found++
		}
	}
	println(found)
}
