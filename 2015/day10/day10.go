package main

import (
	"fmt"
	"strconv"
)

func itoa(n int) []byte {
	return []byte(strconv.Itoa(n))
}

func las(input []byte) []byte {
	result := make([]byte, 0, len(input)*2)

	same := 1
	for a := 1; a < len(input); a += 1 {
		if input[a] == input[a-1] {
			same += 1
		} else {
			result = append(append(result, itoa(same)...), input[a-1])
			same = 1
		}
	}
	// append the last numbers
	return append(append(result, itoa(same)...), input[len(input)-1])
}

func main() {
	//times := 40
	times := 50

	in := []byte("1321131112")
	for i := 0; i < times; i += 1 {
		in = las(in)
	}
	fmt.Println(len(in))
}
