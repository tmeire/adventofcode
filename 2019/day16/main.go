package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

var root = []int{0, 1, 0, -1}

func multiplier(x, y int) int {
	return root[((y+1)/(x+1))%4]
}

func main() {
	bdata, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err.Error())
	}

	offset, err := strconv.Atoi(string(bdata[0:7]))
	if err != nil {
		panic(err.Error())
	}
	println(offset)

	data := make([]int, len(bdata))
	for i := range data {
		data[i] = int(bdata[i]) - 48
	}
	fmt.Printf("%#v\n", data[0:8])

	newdata := make([]int, len(data))
	for phase := 0; phase < 100; phase++ {
		for i := 0; i < len(data); i++ {
			var v int
			for j, x := range data {
				v += int(x) * multiplier(i, j)
			}
			newdata[i] = int(abs(v) % 10)
		}

		data, newdata = newdata, data
	}

	fmt.Printf("%#v\n", data[0:8])

	data = make([]int, len(bdata)*10000)
	for i := range data {
		data[i] = int(bdata[i%len(bdata)]) - 48
	}
	for phase := 0; phase < 100; phase++ {
		for repeat := len(data) - 2; repeat > offset-5; repeat-- {
			n := data[repeat+1] + data[repeat]
			data[repeat] = abs(n) % 10
		}
	}
	fmt.Printf("%#v\n", data[offset:offset+8])

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
