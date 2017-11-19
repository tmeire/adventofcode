package main

import (
	"fmt"
	"testing"
)

func TestOperations(t *testing.T) {

	fixtures := []struct {
		m   Operation
		in  string
		out string
	}{
		{SwapPos{0, 1}, "dhcgebfa", "hdcgebfa"},

		{SwapLetter{[]byte{'d'}, []byte{'b'}}, "ebcda", "edcba"},

		{RotateX{false, 7}, "afdhbgce", "fdhbgcea"},

		{RotateLetter{[]byte{'b'}}, "abdec", "ecabd"},
		{RotateLetter{[]byte{'c'}}, "edcbgafh", "afhedcbg"},
		{RotateLetter{[]byte{0x62}}, "abhgcfde", "deabhgcf"},
		{RotateLetter{[]byte{0x61}}, "dhbcagfe", "bcagfedh"},
		{RotateLetter{[]byte{0x68}}, "cbdagfeh", "hcbdagfe"},

		{Move{1, 4}, "bcdea", "bdeac"},
		{Move{4, 1}, "bcdea", "bacde"},

		{Reverse{0, 4}, "edcba", "abcde"},
		{Reverse{4, 7}, "hcbdagfe", "hcbdefga"},
	}

	for idx := 0; idx < 8; idx++ {
		opt := 0
		if idx >= 4 {
			opt = 1
		}
		rot := (1 + idx + opt)
		fmt.Println(idx, " + ", rot, "->", idx+rot, (idx+rot)%8)
	}

	for _, f := range fixtures {
		r := f.m.execute([]byte(f.in))

		if string(r) != f.out {
			t.Fatalf("Expected '" + f.out + "', got " + string(r))
		}

		u := f.m.undo([]byte(f.out))

		if string(u) != f.in {
			t.Fatalf("Expected '" + f.in + "', got " + string(u))
		}
	}
}
