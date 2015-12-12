package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile("[-]?[0-9]+")

func main() {
	t, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read input")
	}
	sum := 0
	for _, s := range regex.FindAll(t, -1) {
		i, _ := strconv.Atoi(string(s))
		sum += i
	}
	fmt.Printf("sum: %d\n", sum)
}
