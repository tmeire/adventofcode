package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func hash(h string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(h)))
}

func repeated(s string) byte {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return s[i]
		}
	}
	return '-'
}

func hasher(salt string) func(int) string {
	hashes := make(map[int]string)

	return func(i int) string {
		v := fmt.Sprintf("%s%d", salt, i)

		h, ok := hashes[i]
		if !ok {
			h = hash(v)

			for j := 0; j < 2016; j++ {
				h = hash(h)
			}

			hashes[i] = h
		}
		return h
	}
}

func main() {
	salt := "jlmsuwbz"

	md5 := hasher(salt)

	keys := 0
	for i := 0; keys < 64; i++ {
		ho := md5(i)
		r := repeated(ho)
		if r == '-' {
			continue
		}

		pattern := string([]byte{r, r, r, r, r})
		for j := 1; j <= 1000; j++ {
			h := md5(i + j)
			if strings.Contains(h, pattern) {
				keys++
				fmt.Println(i, keys, ho, h)
				break
			}
		}
	}
}
