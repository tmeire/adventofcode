package main

import "fmt"

func dragon(in []byte) []byte {
	l := len(in)
	in = append(in, 0)
	for i := 0; i < l; i++ {
		var x byte
		if in[l-1-i] == 0 {
			x = 1
		} else {
			x = 0
		}
		in = append(in, x)
	}
	return in
}

func checksum(in []byte) []byte {
	out := make([]byte, 0, len(in)/2)
	for i := 0; i < len(in); i += 2 {
		x := byte(0)
		if in[i] == in[i+1] {
			x = byte(1)
		}
		out = append(out, x)
	}
	return out
}

func fill(start []byte, length int) []byte {
	b := make([]byte, 0, length)
	b = append(b, start...)
	for len(b) < length {
		b = dragon(b)
	}
	b = b[:length]

	c := checksum(b)
	for len(c)%2 == 0 {
		c = checksum(c)
	}
	return c
}

func toString(b []byte) string {
	o := make([]byte, len(b))
	for i, bb := range b {
		if bb == 0 {
			o[i] = '0'
		} else {
			o[i] = '1'
		}
	}
	return string(o)
}

func main() {
	s := []byte("10001001100000001")
	start := make([]byte, len(s))
	for i, ss := range s {
		if ss == '1' {
			start[i] = 1
		} else {
			start[i] = 0
		}
	}

	cks := fill(start, 272)
	fmt.Println(toString(cks))

	cks = fill(start, 35651584)
	fmt.Println(toString(cks))
}
