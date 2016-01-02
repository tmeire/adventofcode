package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	key := "yzbqklnj"

	i := 0
	for {
		v := fmt.Sprintf("%s%d", key, i)

		h := md5.New()
		io.WriteString(h, v)

		r := fmt.Sprintf("%x", h.Sum(nil))
		//if r[0:5] == "00000" {
		if r[0:6] == "000000" {
			fmt.Printf("Lowest number is %d, hash was '%s'\n", i, r)
			return
		}
		i += 1
	}
}
